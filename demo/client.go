package demo

import (
	"context"

	"github.com/gs02048/micro/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "qingke.demo"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DemoClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), "etcd://default/"+AppID)
	if err != nil {
		return nil, err
	}
	return NewDemoClient(cc), nil
}
