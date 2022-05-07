package server

import (
	"context"
	"fmt"
	"github.com/journeymidnight/log"
	"google.golang.org/grpc"
	"time"
	"twins/proto"
)

type TwinsRpcClient struct {
	rpcClient proto.TwinsClient
	logger    log.Logger
}

func NewTwinsClient(address string, dialTimeout time.Duration, logger log.Logger) (*TwinsRpcClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	cancel()
	if err != nil {
		return nil, fmt.Errorf("dial %s: %w", address, err)
	}

	return &TwinsRpcClient{
		rpcClient: proto.NewTwinsClient(conn),
		logger:    logger,
	}, nil
}

func (c *TwinsRpcClient) SendHeartBeat(ctx context.Context) (*proto.ResponseHeartbeat, error) {
	return c.rpcClient.Heartbeat(ctx, &proto.RequestHeartbeat{
		Timestamp: time.Now().UTC().UnixNano(),
	})
}
