package p2pconfig

import "time"

const (
	ScoreOptionKey              = "internal-peer-scoring"
	AppScoreWeightKey           = "app-specific-score-weight"
	MaxDebugLogsKey             = "max-debug-logs"
	ScoreOptionDecayIntervalKey = "decay-interval"
	ScoreOptionDecayToZeroKey   = "decay-to-zero"
	ScoreOptionPenaltiesKey     = "penalties"
	ScoreOptionRewardsKey       = "rewards"
	ScoreOptionThresholdsKey    = "thresholds"
	ScoreOptionBehaviourKey     = "behaviour"
	ScoreOptionTopicKey         = "topic"
)

// InternalPeerScoring gossipsub scoring option configuration parameters.
type InternalPeerScoring struct {
	// AppSpecificScoreWeight is the  weight for app-specific scores. It is used to scale the app-specific
	// scores to the same range as the other scores.
	AppSpecificScoreWeight float64 `validate:"gt=0,lte=1" mapstructure:"app-specific-score-weight"`
	// MaxDebugLogs sets the max number of debug/trace log events per second. Logs emitted above
	// this threshold are dropped.
	MaxDebugLogs uint32 `validate:"lte=50" mapstructure:"max-debug-logs"`
	// DecayInterval is the  decay interval for the overall score of a peer at the GossipSub scoring
	// system.
	DecayInterval time.Duration `validate:"gte=1m" mapstructure:"decay-interval"`
	// DecayToZero is the  decay to zero for the overall score of a peer at the GossipSub scoring system.
	// It defines the maximum value below which a peer scoring counter is reset to zero.
	// This is to prevent the counter from decaying to a very small value.
	DecayToZero     float64                    `validate:"required" mapstructure:"decay-to-zero"`
	Penalties       ScoreOptionPenalties       `validate:"required" mapstructure:"penalties"`
	Rewards         ScoreOptionRewards         `validate:"required" mapstructure:"rewards"`
	Thresholds      ScoreOptionThresholds      `validate:"required" mapstructure:"thresholds"`
	Behaviour       ScoreOptionBehaviour       `validate:"required" mapstructure:"behaviour"`
	TopicValidation ScoreOptionTopicValidation `validate:"required" mapstructure:"topic"`
}

const (
	MaxAppSpecificPenaltyKey      = "max-app-specific-penalty"
	MinAppSpecificPenaltyKey      = "min-app-specific-penalty"
	UnknownIdentityPenaltyKey     = "unknown-identity-penalty"
	InvalidSubscriptionPenaltyKey = "invalid-subscription-penalty"
)

// ScoreOptionPenalties score option penalty configuration parameters.
type ScoreOptionPenalties struct {
	// MaxAppSpecificPenalty the maximum penalty for sever offenses that we apply to a remote node score. The score
	// mechanism of GossipSub in Flow is designed in a way that all other infractions are penalized with a fraction of
	// this value.
	MaxAppSpecificPenalty float64 `validate:"lt=0" mapstructure:"max-app-specific-penalty"`
	// MinAppSpecificPenalty the minimum penalty for sever offenses that we apply to a remote node score.
	MinAppSpecificPenalty float64 `validate:"lt=0" mapstructure:"min-app-specific-penalty"`
	// UnknownIdentityPenalty is the  penalty for unknown identity. It is applied to the peer's score when
	// the peer is not in the identity list.
	UnknownIdentityPenalty float64 `validate:"lt=0" mapstructure:"unknown-identity-penalty"`
	// InvalidSubscriptionPenalty is the  penalty for invalid subscription. It is applied to the peer's score when
	// the peer subscribes to a topic that it is not authorized to subscribe to.
	InvalidSubscriptionPenalty float64 `validate:"lt=0" mapstructure:"invalid-subscription-penalty"`
}

const (
	MaxAppSpecificRewardKey = "max-app-specific-reward"
	StakedIdentityRewardKey = "staked-identity-reward"
)

// ScoreOptionRewards score option rewards configuration parameters.
type ScoreOptionRewards struct {
	// MaxAppSpecificReward is the  reward for well-behaving staked peers. If a peer does not have
	// any misbehavior record, e.g., invalid subscription, invalid message, etc., it will be rewarded with this score.
	MaxAppSpecificReward float64 `validate:"gt=0" mapstructure:"max-app-specific-reward"`
	// StakedIdentityReward is the  reward for staking peers. It is applied to the peer's score when
	// the peer does not have any misbehavior record, e.g., invalid subscription, invalid message, etc.
	// The purpose is to reward the staking peers for their contribution to the network and prioritize them in neighbor selection.
	StakedIdentityReward float64 `validate:"gt=0" mapstructure:"staked-identity-reward"`
}

const (
	GossipThresholdKey             = "gossip"
	PublishThresholdKey            = "publish"
	GraylistThresholdKey           = "graylist"
	AcceptPXThresholdKey           = "accept-px"
	OpportunisticGraftThresholdKey = "opportunistic-graft"
)

// ScoreOptionThresholds score option threshold configuration parameters.
type ScoreOptionThresholds struct {
	// Gossip when a peer's penalty drops below this threshold,
	// no gossip is emitted towards that peer and gossip from that peer is ignored.
	Gossip float64 `validate:"lt=0" mapstructure:"gossip"`
	// Publish when a peer's penalty drops below this threshold,
	// self-published messages are not propagated towards this peer.
	Publish float64 `validate:"lt=0" mapstructure:"publish"`
	// Graylist when a peer's penalty drops below this threshold, the peer is graylisted, i.e.,
	// incoming RPCs from the peer are ignored.
	Graylist float64 `validate:"lt=0" mapstructure:"graylist"`
	// AcceptPX when a peer sends us PX information with a prune, we only accept it and connect to the supplied
	// peers if the originating peer's penalty exceeds this threshold.
	AcceptPX float64 `validate:"gt=0" mapstructure:"accept-px"`
	// OpportunisticGraft when the median peer penalty in the mesh drops below this value,
	// the peer may select more peers with penalty above the median to opportunistically graft on the mesh.
	OpportunisticGraft float64 `validate:"gt=0" mapstructure:"opportunistic-graft"`
}

const (
	BehaviourPenaltyThresholdKey = "penalty-threshold"
	BehaviourPenaltyWeightKey    = "penalty-weight"
	BehaviourPenaltyDecayKey     = "penalty-decay"
)

// ScoreOptionBehaviour score option behaviour configuration parameters.
type ScoreOptionBehaviour struct {
	// PenaltyThreshold is the threshold when the behavior of a peer is considered as bad by GossipSub.
	// Currently, the misbehavior is defined as advertising an iHave without responding to the iWants (iHave broken promises), as well as attempting
	// on GRAFT when the peer is considered for a PRUNE backoff, i.e., the local peer does not allow the peer to join the local topic mesh
	// for a while, and the remote peer keep attempting on GRAFT (aka GRAFT flood).
	// When the misbehavior counter of a peer goes beyond this threshold, the peer is penalized by BehaviorPenaltyWeight (see below) for the excess misbehavior.
	//
	// An iHave broken promise means that a peer advertises an iHave for a message, but does not respond to the iWant requests for that message.
	// For iHave broken promises, the gossipsub scoring works as follows:
	// It samples ONLY A SINGLE iHave out of the entire RPC.
	// If that iHave is not followed by an actual message within the next 3 seconds, the peer misbehavior counter is incremented by 1.
	//
	// The counter is also decayed by (0.99) every decay interval (DecayInterval) i.e., every minute.
	// Note that misbehaviors are counted by GossipSub across all topics (and is different from the Application Layer Misbehaviors that we count through
	// the ALSP system).
	PenaltyThreshold float64 `validate:"gt=0" mapstructure:"penalty-threshold"`
	// PenaltyWeight is the weight for applying penalty when a peer misbehavior goes beyond the threshold.
	// Misbehavior of a peer at gossipsub layer is defined as advertising an iHave without responding to the iWants (broken promises), as well as attempting
	// on GRAFT when the peer is considered for a PRUNE backoff, i.e., the local peer does not allow the peer to join the local topic mesh
	// This is detected by the GossipSub scoring system, and the peer is penalized by BehaviorPenaltyWeight.
	//
	// An iHave broken promise means that a peer advertises an iHave for a message, but does not respond to the iWant requests for that message.
	// For iHave broken promises, the gossipsub scoring works as follows:
	// It samples ONLY A SINGLE iHave out of the entire RPC.
	// If that iHave is not followed by an actual message within the next 3 seconds, the peer misbehavior counter is incremented by 1.
	PenaltyWeight float64 `validate:"lt=0" mapstructure:"penalty-weight"`
	// PenaltyDecay is the decay interval for the misbehavior counter of a peer. The misbehavior counter is
	// incremented by GossipSub for iHave broken promises or the GRAFT flooding attacks (i.e., each GRAFT received from a remote peer while that peer is on a PRUNE backoff).
	//
	// An iHave broken promise means that a peer advertises an iHave for a message, but does not respond to the iWant requests for that message.
	// For iHave broken promises, the gossipsub scoring works as follows:
	// It samples ONLY A SINGLE iHave out of the entire RPC.
	// If that iHave is not followed by an actual message within the next 3 seconds, the peer misbehavior counter is incremented by 1.
	// This means that regardless of how many iHave broken promises an RPC contains, the misbehavior counter is incremented by 1.
	// That is why we decay the misbehavior counter very slow, as this counter indicates a severe misbehavior.
	// The misbehavior counter is decayed per decay interval (i.e., DecayInterval = 1 minute) by GossipSub.
	//
	// Note that misbehaviors are counted by GossipSub across all topics (and is different from the Application Layer Misbehaviors that we count through
	// the ALSP system that is based on the engines report).
	PenaltyDecay float64 `validate:"gt=0,lt=1" mapstructure:"penalty-decay"`
}

const (
	SkipAtomicValidationKey           = "skip-atomic-validation"
	InvalidMessageDeliveriesWeightKey = "invalid-message-deliveries-weight"
	InvalidMessageDeliveriesDecayKey  = "invalid-message-deliveries-decay"
	TimeInMeshQuantumKey              = "time-in-mesh-quantum"
	TopicWeightKey                    = "topic-weight"
	MeshMessageDeliveriesDecayKey     = "mesh-message-deliveries-decay"
	MeshMessageDeliveriesCapKey       = "mesh-message-deliveries-cap"
	MeshMessageDeliveryThresholdKey   = "mesh-message-deliveries-threshold"
	MeshDeliveriesWeightKey           = "mesh-deliveries-weight"
	MeshMessageDeliveriesWindowKey    = "mesh-message-deliveries-window"
	MeshMessageDeliveryActivationKey  = "mesh-message-delivery-activation"
)

// ScoreOptionTopicValidation score option topic validation configuration parameters.
type ScoreOptionTopicValidation struct {
	// SkipAtomicValidation is the  value for the skip atomic validation flag for topics.
	// If set it to true, the gossipsub parameter validation will not fail if we leave some of the
	// topic parameters at their  values, i.e., zero.
	SkipAtomicValidation bool `validate:"required" mapstructure:"skip-atomic-validation"`
	// InvalidMessageDeliveriesWeight this value is applied to the square of the number of invalid message deliveries on a topic.
	// It is used to penalize peers that send invalid messages. By an invalid message, we mean a message that is not signed by the
	// publisher, or a message that is not signed by the peer that sent it.
	InvalidMessageDeliveriesWeight float64 `validate:"lt=0" mapstructure:"invalid-message-deliveries-weight"`
	// InvalidMessageDeliveriesDecay decay factor used to decay the number of invalid message deliveries.
	// The total number of invalid message deliveries is multiplied by this factor at each heartbeat interval to
	// decay the number of invalid message deliveries, and prevent the peer from being disconnected if it stops
	// sending invalid messages.
	InvalidMessageDeliveriesDecay float64 `validate:"gt=0,lt=1" mapstructure:"invalid-message-deliveries-decay"`
	// TimeInMeshQuantum is the  time in mesh quantum for the GossipSub scoring system. It is used to gauge
	// a discrete time interval for the time in mesh counter.
	TimeInMeshQuantum time.Duration `validate:"gte=1h" mapstructure:"time-in-mesh-quantum"`
	// Weight is the  weight of a topic in the GossipSub scoring system.
	// The overall score of a peer in a topic mesh is multiplied by the weight of the topic when calculating the overall score of the peer.
	TopicWeight float64 `validate:"gt=0" mapstructure:"topic-weight"`
	// MeshMessageDeliveriesDecay is applied to the number of actual message deliveries in a topic mesh
	// at each decay interval (i.e., DecayInterval).
	// It is used to decay the number of actual message deliveries, and prevents past message
	// deliveries from affecting the current score of the peer.
	MeshMessageDeliveriesDecay float64 `validate:"gt=0" mapstructure:"mesh-message-deliveries-decay"`
	// MeshMessageDeliveriesCap is the maximum number of actual message deliveries in a topic
	// mesh that is used to calculate the score of a peer in that topic mesh.
	MeshMessageDeliveriesCap float64 `validate:"gt=0" mapstructure:"mesh-message-deliveries-cap"`
	// MeshMessageDeliveryThreshold is the threshold for the number of actual message deliveries in a
	// topic mesh that is used to calculate the score of a peer in that topic mesh.
	// If the number of actual message deliveries in a topic mesh is less than this value,
	// the peer will be penalized by square of the difference between the actual message deliveries and the threshold,
	// i.e., -w * (actual - threshold)^2 where `actual` and `threshold` are the actual message deliveries and the
	// threshold, respectively, and `w` is the weight (i.e., MeshMessageDeliveriesWeight).
	MeshMessageDeliveryThreshold float64 `validate:"gt=0" mapstructure:"mesh-message-deliveries-threshold"`
	// MeshDeliveriesWeight is the weight for applying penalty when a peer is under-performing in a topic mesh.
	// Upon every decay interval, if the number of actual message deliveries is less than the topic mesh message deliveries threshold
	// (i.e., MeshMessageDeliveriesThreshold), the peer will be penalized by square of the difference between the actual
	// message deliveries and the threshold, multiplied by this weight, i.e., -w * (actual - threshold)^2 where w is the weight, and
	// `actual` and `threshold` are the actual message deliveries and the threshold, respectively.
	MeshDeliveriesWeight float64 `validate:"lt=0" mapstructure:"mesh-deliveries-weight"`
	// MeshMessageDeliveriesWindow is the window size is time interval that we count a delivery of an already
	// seen message towards the score of a peer in a topic mesh. The delivery is counted
	// by GossipSub only if the previous sender of the message is different from the current sender.
	MeshMessageDeliveriesWindow time.Duration `validate:"gte=1m" mapstructure:"mesh-message-deliveries-window"`
	// MeshMessageDeliveryActivation is the time interval that we wait for a new peer that joins a topic mesh
	// till start counting the number of actual message deliveries of that peer in that topic mesh.
	MeshMessageDeliveryActivation time.Duration `validate:"gte=2m" mapstructure:"mesh-message-delivery-activation"`
}
