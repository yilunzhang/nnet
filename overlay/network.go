package overlay

import (
	"time"

	"github.com/nknorg/nnet/node"
	"github.com/nknorg/nnet/overlay/routing"
	"github.com/nknorg/nnet/protobuf"
)

// Network is the overlay network interface
type Network interface {
	Start(isCreate bool) error
	Stop(error)
	Join(seedNodeAddr string) error
	GetLocalNode() *node.LocalNode
	GetRouters() []routing.Router
	ApplyMiddleware(interface{}) error
	SendMessageAsync(msg *protobuf.Message, routingType protobuf.RoutingType) (success bool, err error)
	SendMessageSync(msg *protobuf.Message, routingType protobuf.RoutingType, replyTimeout time.Duration) (reply *protobuf.Message, success bool, err error)
}
