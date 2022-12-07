package corruptlibp2p

import (
	pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/require"
	pubsub "github.com/yhassanzadeh13/go-libp2p-pubsub"
	"testing"
)

type ControlMessage int

// SpammerGossipSub is a wrapper around the GossipSubRouter that allows us to
// spam the victim with junk control messages.
type SpammerGossipSub struct {
	router *pubsub.GossipSubRouter
}

func NewSpammerGossipSubRouter(router *pubsub.GossipSubRouter) *SpammerGossipSub {
	return &SpammerGossipSub{
		router: router,
	}
}

// SpamIHave spams the victim with junk iHave messages.
// msgCount is the number of iHave messages to send.
// msgSize is the number of messageIDs to include in each iHave message.
func (s *SpammerGossipSub) SpamIHave(victim peer.ID, ctlMessages []pb.ControlMessage) {
	for _, ctlMessage := range ctlMessages {
		s.router.SendControl(victim, &ctlMessage)
	}
}

// GenerateIHaveCtlMessages generates IHAVE control messages before they are sent so the test can prepare
// to receive them before they are sent by the spammer.
func (s *SpammerGossipSub) GenerateIHaveCtlMessages(t *testing.T, msgCount, msgSize int) []pb.ControlMessage {
	//var ctlMessageMap = make(map[string]pb.ControlMessage)
	var iHaveCtlMsgs []pb.ControlMessage
	for i := 0; i < msgCount; i++ {
		iHaveCtlMsg := GossipSubCtrlFixture(WithIHave(msgCount, msgSize))

		iHaves := iHaveCtlMsg.GetIhave()
		require.Equal(t, msgCount, len(iHaves))
		iHaveCtlMsgs = append(iHaveCtlMsgs, *iHaveCtlMsg)
	}
	return iHaveCtlMsgs
}

// TODO: SpamIWant, SpamGraft, SpamPrune.
