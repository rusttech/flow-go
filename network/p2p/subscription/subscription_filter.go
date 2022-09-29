package subscription

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/id"
	"github.com/onflow/flow-go/network/channels"
)

// RoleBasedFilter implements a subscription filter that filters subscriptions based on a node's role.
type RoleBasedFilter struct {
	idProvider id.IdentityProvider
	myRole     flow.Role
}

const UnstakedRole = flow.Role(0)

var _ pubsub.SubscriptionFilter = (*RoleBasedFilter)(nil)

func NewRoleBasedFilter(role flow.Role, idProvider id.IdentityProvider) *RoleBasedFilter {
	filter := &RoleBasedFilter{
		idProvider: idProvider,
		myRole:     role,
	}

	return filter
}

func (f *RoleBasedFilter) getRole(pid peer.ID) flow.Role {
	if flowId, ok := f.idProvider.ByPeerID(pid); ok {
		return flowId.Role
	}

	return UnstakedRole
}

func (f *RoleBasedFilter) allowed(role flow.Role, topic string) bool {
	channel, ok := channels.ChannelFromTopic(channels.Topic(topic))
	if !ok {
		return false
	}

	if !role.Valid() {
		// TODO: eventually we should have block proposals relayed on a separate
		// channel on the public network. For now, we need to make sure that
		// full observer nodes can subscribe to the block proposal channel.
		return append(channels.PublicChannels(), channels.ReceiveBlocks).Contains(channel)
	} else {
		if roles, ok := channels.RolesByChannel(channel); ok {
			return roles.Contains(role)
		}

		return false
	}
}

func (f *RoleBasedFilter) CanSubscribe(topic string) bool {
	return f.allowed(f.myRole, topic)
}

func (f *RoleBasedFilter) FilterIncomingSubscriptions(from peer.ID, opts []*pb.RPC_SubOpts) ([]*pb.RPC_SubOpts, error) {
	role := f.getRole(from)
	var filtered []*pb.RPC_SubOpts

	for _, opt := range opts {
		if f.allowed(role, opt.GetTopicid()) {
			filtered = append(filtered, opt)
		}
	}

	return filtered, nil
}