package scoring

import (
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/network/p2p"
	netcache "github.com/onflow/flow-go/network/p2p/cache"
	p2pmsg "github.com/onflow/flow-go/network/p2p/message"
	"github.com/onflow/flow-go/network/p2p/p2plogging"
	"github.com/onflow/flow-go/utils/logging"
	"github.com/onflow/flow-go/utils/rand"
)

const (
	// MaxDecay is the maximum decay value for the application specific penalty.
	// Spam record will be initialized with a decay value between .5 , .7 and this value will then be decayed up to .99 on consecutive misbehavior's,
	// The maximum decay value decays the penalty by 1% every second.
	// assume:
	//     penalty = -100 (the maximum application specific penalty is -100)
	//     skipDecayThreshold = -0.1
	// it takes around 459 seconds for the penalty to decay to reach greater than -0.1 and turn into 0.
	//     x * 0.99 ^ n > -0.1 (assuming negative x).
	//     0.99 ^ n > -0.1 / x
	// Now we can take the logarithm of both sides (with any base, but let's use base 10 for simplicity).
	//     log( 0.99 ^ n ) < log( 0.1 / x )
	// Using the properties of logarithms, we can bring down the exponent:
	//     n * log( 0.99 ) < log( -0.1 / x )
	// And finally, we can solve for n:
	//     n > log( -0.1 / x ) / log( 0.99 )
	// We can plug in x = -100:
	//     n > log( -0.1 / -100 ) / log( 0.99 )
	//     n > log( 0.001 ) / log( 0.99 )
	//     n > -3 / log( 0.99 )
	//     n >  458.22
        // MinSpamPenaltyDecaySpeed is minimum the speed at which the spam penalty value of a peer is decayed. 
        // The decay is applied geometrically, i.e., `newPenalty = oldPenalty * decay`, hence, the higher decay value 
        // indicates a lower decay speed, i.e., it takes more number of heartbeat intervals to decay a penalty back to 
        // zero when the decay value is high.  
	MinSpamPenaltyDecaySpeed = 0.99
	// skipDecayThreshold is the threshold for which when the negative penalty is above this value, the decay function will not be called.
	// instead, the penalty will be set to 0. This is to prevent the penalty from keeping a small negative value for a long time.
	skipDecayThreshold = -0.1
	// graftMisbehaviourPenalty is the penalty applied to the application specific penalty when a peer conducts a graft misbehaviour.
	graftMisbehaviourPenalty = -10
	// pruneMisbehaviourPenalty is the penalty applied to the application specific penalty when a peer conducts a prune misbehaviour.
	pruneMisbehaviourPenalty = -10
	// iHaveMisbehaviourPenalty is the penalty applied to the application specific penalty when a peer conducts a iHave misbehaviour.
	iHaveMisbehaviourPenalty = -10
	// iWantMisbehaviourPenalty is the penalty applied to the application specific penalty when a peer conducts a iWant misbehaviour.
	iWantMisbehaviourPenalty = -10
	// rpcPublishMessageMisbehaviourPenalty is the penalty applied to the application specific penalty when a peer conducts a RpcPublishMessageMisbehaviourPenalty misbehaviour.
	rpcPublishMessageMisbehaviourPenalty = -10
	// randomDecayFloatDensity used to generate a secure random Uint64 that is used to generate the random initial decay for app score records.
	randomDecayFloatDensity = uint64(1000)
)

// GossipSubCtrlMsgPenaltyValue is the penalty value for each control message type.
type GossipSubCtrlMsgPenaltyValue struct {
	Graft             float64 // penalty value for an individual graft message misbehaviour.
	Prune             float64 // penalty value for an individual prune message misbehaviour.
	IHave             float64 // penalty value for an individual iHave message misbehaviour.
	IWant             float64 // penalty value for an individual iWant message misbehaviour.
	RpcPublishMessage float64 // penalty value for an individual RpcPublishMessage message misbehaviour.
}

// DefaultGossipSubCtrlMsgPenaltyValue returns the default penalty value for each control message type.
func DefaultGossipSubCtrlMsgPenaltyValue() GossipSubCtrlMsgPenaltyValue {
	return GossipSubCtrlMsgPenaltyValue{
		Graft:             graftMisbehaviourPenalty,
		Prune:             pruneMisbehaviourPenalty,
		IHave:             iHaveMisbehaviourPenalty,
		IWant:             iWantMisbehaviourPenalty,
		RpcPublishMessage: rpcPublishMessageMisbehaviourPenalty,
	}
}

// GossipSubAppSpecificScoreRegistry is the registry for the application specific score of peers in the GossipSub protocol.
// The application specific score is part of the overall score of a peer, and is used to determine the peer's score based
// on its behavior related to the application (Flow protocol).
// This registry holds the view of the local peer of the application specific score of other peers in the network based
// on what it has observed from the network.
// Similar to the GossipSub score, the application specific score is meant to be private to the local peer, and is not
// shared with other peers in the network.
type GossipSubAppSpecificScoreRegistry struct {
	logger     zerolog.Logger
	idProvider module.IdentityProvider
	// spamScoreCache currently only holds the control message misbehaviour penalty (spam related penalty).
	spamScoreCache p2p.GossipSubSpamRecordCache
	penalty        GossipSubCtrlMsgPenaltyValue
	// initial application specific penalty record, used to initialize the penalty cache entry.
	init                      func() (p2p.GossipSubSpamRecord, error)
	validator                 p2p.SubscriptionValidator
	initDecayLowerBound       float64
	initDecayUpperBound       float64
	increaseDecayThreshold    float64
	decayThresholdIncrementer float64
}

// GossipSubAppSpecificScoreRegistryParams is the configuration for the GossipSubAppSpecificScoreRegistry.
// The configuration is used to initialize the registry.
type GossipSubAppSpecificScoreRegistryParams struct {
	Logger zerolog.Logger

	// Validator is the subscription validator used to validate the subscriptions of peers, and determine if a peer is
	// authorized to subscribe to a topic.
	Validator p2p.SubscriptionValidator

	// Penalty encapsulates the penalty unit for each control message type misbehaviour.
	Penalty GossipSubCtrlMsgPenaltyValue

	// IdProvider is the identity provider used to translate peer ids at the networking layer to Flow identifiers (if
	// an authorized peer is found).
	IdProvider module.IdentityProvider

	// Init is a factory function that returns a new GossipSubSpamRecord. It is used to initialize the spam record of
	// a peer when the peer is first observed by the local peer.
	Init func() (p2p.GossipSubSpamRecord, error)

	// CacheFactory is a factory function that returns a new GossipSubSpamRecordCache. It is used to initialize the spamScoreCache.
	// The cache is used to store the application specific penalty of peers.
	CacheFactory func() p2p.GossipSubSpamRecordCache

	// InitDecayLowerBound is the lower bound on the decay value for a spam record when initialized.
	// A random value in a range of InitDecayLowerBound and InitDecayUpperBound is used when initializing the decay
	// of a spam record.
	InitDecayLowerBound float64

	// InitDecayUpperBound is the upper bound on the decay value for a spam record when initialized.
	InitDecayUpperBound float64

// SlowerDecayPenaltyThreshold defines the penalty level which the decay rate is reduced by `DecayRateDecremen` every time the penalty of a node falls below the threshold, thereby slowing down the decay process. This mechanism ensures that malicious nodes experience longer decay periods, while honest nodes benefit from quicker decay.
SlowerDecayPenaltyThreshold float64

// DecayRateDecrement defines the value by which the decay rate is decreased every time the penalty is below the SlowerDecayPenaltyThreshold. A reduced decay rate extends the time it takes for penalties to diminish.
DecayRateDecrement float64
}

// NewGossipSubAppSpecificScoreRegistry returns a new GossipSubAppSpecificScoreRegistry.
// Args:
//
//	params: the parameters for the registry.
//
// Returns:
//
//	a new GossipSubAppSpecificScoreRegistry.
func NewGossipSubAppSpecificScoreRegistry(params *GossipSubAppSpecificScoreRegistryParams) *GossipSubAppSpecificScoreRegistry {
	reg := &GossipSubAppSpecificScoreRegistry{
		logger:                    params.Logger.With().Str("module", "app_score_registry").Logger(),
		spamScoreCache:            params.CacheFactory(),
		penalty:                   params.Penalty,
		init:                      params.Init,
		validator:                 params.Validator,
		idProvider:                params.IdProvider,
		initDecayLowerBound:       params.InitDecayLowerBound,
		initDecayUpperBound:       params.InitDecayUpperBound,
		increaseDecayThreshold:    params.IncreaseDecayThreshold,
		decayThresholdIncrementer: params.DecayThresholdIncrementer,
	}

	return reg
}

var _ p2p.GossipSubInvCtrlMsgNotifConsumer = (*GossipSubAppSpecificScoreRegistry)(nil)

// AppSpecificScoreFunc returns the application specific penalty function that is called by the GossipSub protocol to determine the application specific penalty of a peer.
func (r *GossipSubAppSpecificScoreRegistry) AppSpecificScoreFunc() func(peer.ID) float64 {
	return func(pid peer.ID) float64 {
		appSpecificScore := float64(0)

		lg := r.logger.With().Str("peer_id", p2plogging.PeerId(pid)).Logger()
		// (1) spam penalty: the penalty is applied to the application specific penalty when a peer conducts a spamming misbehaviour.
		spamRecord, err, spamRecordExists := r.spamScoreCache.Get(pid)
		if err != nil {
			// the error is considered fatal as it means the cache is not working properly.
			// we should not continue with the execution as it may lead to routing attack vulnerability.
			r.logger.Fatal().Str("peer_id", p2plogging.PeerId(pid)).Err(err).Msg("could not get application specific penalty for peer")
			return appSpecificScore // unreachable, but added to avoid proceeding with the execution if log level is changed.
		}

		if spamRecordExists {
			lg = lg.With().Float64("spam_penalty", spamRecord.Penalty).Logger()
			appSpecificScore += spamRecord.Penalty
		}

		// (2) staking score: for staked peers, a default positive reward is applied only if the peer has no penalty on spamming and subscription.
		// for unknown peers a negative penalty is applied.
		stakingScore, flowId, role := r.stakingScore(pid)
		if stakingScore < 0 {
			lg = lg.With().Float64("staking_penalty", stakingScore).Logger()
			// staking penalty is applied right away.
			appSpecificScore += stakingScore
		}

		if stakingScore >= 0 {
			// (3) subscription penalty: the subscription penalty is applied to the application specific penalty when a
			// peer is subscribed to a topic that it is not allowed to subscribe to based on its role.
			// Note: subscription penalty can be considered only for staked peers, for non-staked peers, we cannot
			// determine the role of the peer.
			subscriptionPenalty := r.subscriptionPenalty(pid, flowId, role)
			lg = lg.With().Float64("subscription_penalty", subscriptionPenalty).Logger()
			if subscriptionPenalty < 0 {
				appSpecificScore += subscriptionPenalty
			}
		}

		// (4) staking reward: for staked peers, a default positive reward is applied only if the peer has no penalty on spamming and subscription.
		if stakingScore > 0 && appSpecificScore == float64(0) {
			lg = lg.With().Float64("staking_reward", stakingScore).Logger()
			appSpecificScore += stakingScore
		}

		lg.Trace().
			Float64("total_app_specific_score", appSpecificScore).
			Msg("application specific penalty computed")

		return appSpecificScore
	}
}

func (r *GossipSubAppSpecificScoreRegistry) stakingScore(pid peer.ID) (float64, flow.Identifier, flow.Role) {
	lg := r.logger.With().Str("peer_id", p2plogging.PeerId(pid)).Logger()

	// checks if peer has a valid Flow protocol identity.
	flowId, err := HasValidFlowIdentity(r.idProvider, pid)
	if err != nil {
		lg.Error().
			Err(err).
			Bool(logging.KeySuspicious, true).
			Msg("invalid peer identity, penalizing peer")
		return DefaultUnknownIdentityPenalty, flow.Identifier{}, 0
	}

	lg = lg.With().
		Hex("flow_id", logging.ID(flowId.NodeID)).
		Str("role", flowId.Role.String()).
		Logger()

	// checks if peer is an access node, and if so, pushes it to the
	// edges of the network by giving the minimum penalty.
	if flowId.Role == flow.RoleAccess {
		lg.Trace().
			Msg("pushing access node to edge by penalizing with minimum penalty value")
		return MinAppSpecificPenalty, flowId.NodeID, flowId.Role
	}

	lg.Trace().
		Msg("rewarding well-behaved non-access node peer with maximum reward value")

	return DefaultStakedIdentityReward, flowId.NodeID, flowId.Role
}

func (r *GossipSubAppSpecificScoreRegistry) subscriptionPenalty(pid peer.ID, flowId flow.Identifier, role flow.Role) float64 {
	// checks if peer has any subscription violation.
	if err := r.validator.CheckSubscribedToAllowedTopics(pid, role); err != nil {
		r.logger.Err(err).
			Str("peer_id", p2plogging.PeerId(pid)).
			Hex("flow_id", logging.ID(flowId)).
			Bool(logging.KeySuspicious, true).
			Msg("invalid subscription detected, penalizing peer")
		return DefaultInvalidSubscriptionPenalty
	}

	return 0
}

// OnInvalidControlMessageNotification is called when a new invalid control message notification is distributed.
// Any error on consuming event must handle internally.
// The implementation must be concurrency safe, but can be blocking.
func (r *GossipSubAppSpecificScoreRegistry) OnInvalidControlMessageNotification(notification *p2p.InvCtrlMsgNotif) {
	// we use mutex to ensure the method is concurrency safe.

	lg := r.logger.With().
		Err(notification.Error).
		Str("peer_id", p2plogging.PeerId(notification.PeerID)).
		Str("misbehavior_type", notification.MsgType.String()).Logger()

	// try initializing the application specific penalty for the peer if it is not yet initialized.
	// this is done to avoid the case where the peer is not yet cached and the application specific penalty is not yet initialized.
	// initialization is successful only if the peer is not yet cached. If any error is occurred during initialization we log a fatal error
	initRecord, err := r.init()
	if err != nil {
		// the error is considered fatal as it means that we cannot generate a random initial decay.
		lg.Fatal().Str("misbehavior_type", notification.MsgType.String()).Msg("failed to initialize app specific score record")
	}
	initialized := r.spamScoreCache.Add(notification.PeerID, initRecord)
	if initialized {
		lg.Trace().Str("peer_id", p2plogging.PeerId(notification.PeerID)).Msg("application specific penalty initialized for peer")
	}

	record, err := r.spamScoreCache.Update(notification.PeerID, func(record p2p.GossipSubSpamRecord) p2p.GossipSubSpamRecord {
		switch notification.MsgType {
		case p2pmsg.CtrlMsgGraft:
			record.Penalty += r.penalty.Graft
		case p2pmsg.CtrlMsgPrune:
			record.Penalty += r.penalty.Prune
		case p2pmsg.CtrlMsgIHave:
			record.Penalty += r.penalty.IHave
		case p2pmsg.CtrlMsgIWant:
			record.Penalty += r.penalty.IWant
		case p2pmsg.RpcPublishMessage:
			record.Penalty += r.penalty.RpcPublishMessage
		default:
			// the error is considered fatal as it means that we have an unsupported misbehaviour type, we should crash the node to prevent routing attack vulnerability.
			lg.Fatal().Str("misbehavior_type", notification.MsgType.String()).Msg("unknown misbehaviour type")
		}
		return record
	})
	if err != nil {
		// any returned error from adjust is non-recoverable and fatal, we crash the node.
		lg.Fatal().Err(err).Msg("could not adjust application specific penalty for peer")
	}

	lg.Debug().
		Float64("app_specific_score", record.Penalty).
		Msg("applied misbehaviour penalty and updated application specific penalty")
}

// DefaultDecayAdjustmentFunc will adjust the decay for a spam record if the record penalty has reached the IncreaseDecayThreshold.
func DefaultDecayAdjustmentFunc(increaseDecayThreshold float64, decayThresholdIncrementer float64) netcache.PreprocessorFunc {
	return func(record p2p.GossipSubSpamRecord, lastUpdated time.Time) (p2p.GossipSubSpamRecord, error) {
		if record.Penalty <= increaseDecayThreshold {
			decay := record.Decay + decayThresholdIncrementer
			if decay > MaxDecay {
				record.Decay = MaxDecay
			} else {
				record.Decay = decay
			}
		}
		return record, nil
	}
}

// DefaultDecayFunction is the default decay function that is used to decay the application specific penalty of a peer.
// It is used if no decay function is provided in the configuration.
// It decays the application specific penalty of a peer if it is negative.
func DefaultDecayFunction() netcache.PreprocessorFunc {
	return func(record p2p.GossipSubSpamRecord, lastUpdated time.Time) (p2p.GossipSubSpamRecord, error) {
		if record.Penalty >= 0 {
			// no need to decay the penalty if it is positive, the reason is currently the app specific penalty
			// is only used to penalize peers. Hence, when there is no reward, there is no need to decay the positive penalty, as
			// no node can accumulate a positive penalty.
			return record, nil
		}

		if record.Penalty > skipDecayThreshold {
			// penalty is negative but greater than the threshold, we set it to 0.
			record.Penalty = 0
			return record, nil
		}

		// penalty is negative and below the threshold, we decay it.
		penalty, err := GeometricDecay(record.Penalty, record.Decay, lastUpdated)
		if err != nil {
			return record, fmt.Errorf("could not decay application specific penalty: %w", err)
		}
		record.Penalty = penalty
		return record, nil
	}
}

// InitAppScoreRecordStateFunc returns a func that will initialize the spac record state for a peer. The initial decay for the
// spam record will be a random value between the lower and upper bounds provided.
// Args:
//   - decayLowerBound: the lower bound of the range for the initial random decay value
//   - decayUpperBound: the upper bound of the range for the initial random decay value
//
// Returns:
//   - callback func that initializes and returns a spam record.
func InitAppScoreRecordStateFunc(decayLowerBound, decayUpperBound float64) func() (p2p.GossipSubSpamRecord, error) {
	return func() (p2p.GossipSubSpamRecord, error) {
		decay, err := initialDecay(decayLowerBound, decayUpperBound)
		if err != nil {
			return p2p.GossipSubSpamRecord{}, err
		}
		return p2p.GossipSubSpamRecord{
			Decay:   decay,
			Penalty: 0,
		}, nil
	}
}

// initialDecay returns a secure random value between decayLowerBound and decayUpperBound.
// Args:
//   - decayLowerBound: the lower bound of the range for the initial random decay value
//   - decayUpperBound: the upper bound of the range for the initial random decay value
//
// Returns:
//   - float64: the random decay value
//   - error: if any error occurs generating the secure random value
//
// All errors returned from this func are considered irrecoverable.
func initialDecay(decayLowerBound, decayUpperBound float64) (float64, error) {
	randomUint64, err := rand.Uint64n(randomDecayFloatDensity)
	if err != nil {
		return 0, fmt.Errorf("failed to generate random float between %f and %f", decayLowerBound, decayUpperBound)
	}
	randomFloat := float64(randomUint64) / float64(randomDecayFloatDensity)
	return randomFloat*(decayUpperBound-decayLowerBound) + decayLowerBound, nil
}
