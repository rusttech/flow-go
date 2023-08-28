package p2p

import (
	"github.com/hashicorp/go-multierror"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/onflow/flow-go/module/component"
	p2pmsg "github.com/onflow/flow-go/network/p2p/message"
)

// GossipSubInspectorNotifDistributor is the interface for the distributor that distributes gossip sub inspector notifications.
// It is used to distribute notifications to the consumers in an asynchronous manner and non-blocking manner.
// The implementation should guarantee that all registered consumers are called upon distribution of a new event.
type GossipSubInspectorNotifDistributor interface {
	component.Component
	// Distribute distributes the event to all the consumers.
	// Any error returned by the distributor is non-recoverable and will cause the node to crash.
	// Implementation must be concurrency safe, and non-blocking.
	Distribute(notification *InvCtrlMsgNotif) error

	// AddConsumer adds a consumer to the distributor. The consumer will be called the distributor distributes a new event.
	// AddConsumer must be concurrency safe. Once a consumer is added, it must be called for all future events.
	// There is no guarantee that the consumer will be called for events that were already received by the distributor.
	AddConsumer(GossipSubInvCtrlMsgNotifConsumer)
}

// ControlMessageTypeErrs map of control message type => error used to track all errors encountered during a single RPC inspection.
type ControlMessageTypeErrs map[p2pmsg.ControlMessageType]error

// Error returns all errors in a single error.
func (m ControlMessageTypeErrs) Error() error {
	var errs *multierror.Error
	for _, err := range m {
		errs = multierror.Append(errs, err)
	}
	return errs.ErrorOrNil()
}

// InvCtrlMsgNotif is the notification sent to the consumer when an invalid control message is received.
// It models the information that is available to the consumer about a misbehaving peer.
type InvCtrlMsgNotif struct {
	// PeerID is the ID of the peer that sent the invalid control message.
	PeerID peer.ID
	// errs map of control message type -> error encountered during validation. Validation is immediately halted for a control message type
	// when an error is encountered, thus we can expect a single error mapped to each control message type if one was encountered.
	errs ControlMessageTypeErrs
}

// Errors returns the notification errors
func (notif *InvCtrlMsgNotif) Errors() ControlMessageTypeErrs {
	return notif.errs
}

// NewInvalidControlMessageNotification returns a new *InvCtrlMsgNotif
func NewInvalidControlMessageNotification(peerID peer.ID, errs ControlMessageTypeErrs) *InvCtrlMsgNotif {
	return &InvCtrlMsgNotif{
		PeerID: peerID,
		errs:   errs,
	}
}

// GossipSubInvCtrlMsgNotifConsumer is the interface for the consumer that consumes gossipsub inspector notifications.
// It is used to consume notifications in an asynchronous manner.
// The implementation must be concurrency safe, but can be blocking. This is due to the fact that the consumer is called
// asynchronously by the distributor.
type GossipSubInvCtrlMsgNotifConsumer interface {
	// OnInvalidControlMessageNotification is called when a new invalid control message notification is distributed.
	// Any error on consuming event must handle internally.
	// The implementation must be concurrency safe, but can be blocking.
	OnInvalidControlMessageNotification(*InvCtrlMsgNotif)
}

// GossipSubInspectorSuite is the interface for the GossipSub inspector suite.
// It encapsulates the rpc inspectors and the notification distributors.
type GossipSubInspectorSuite interface {
	component.Component
	CollectionClusterChangesConsumer
	// InspectFunc returns the inspect function that is used to inspect the gossipsub rpc messages.
	// This function follows a dependency injection pattern, where the inspect function is injected into the gossipsu, and
	// is called whenever a gossipsub rpc message is received.
	InspectFunc() func(peer.ID, *pubsub.RPC) error

	// AddInvalidControlMessageConsumer AddInvCtrlMsgNotifConsumer adds a consumer to the invalid control message notification distributor.
	// This consumer is notified when a misbehaving peer regarding gossipsub control messages is detected. This follows a pub/sub
	// pattern where the consumer is notified when a new notification is published.
	// A consumer is only notified once for each notification, and only receives notifications that were published after it was added.
	AddInvalidControlMessageConsumer(GossipSubInvCtrlMsgNotifConsumer)
}
