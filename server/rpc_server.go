package server

import (
	"context"
	"github.com/journeymidnight/log"
	"google.golang.org/grpc"
	"net"
	"twins/proto"
)

type TwinsRpcServer struct {
	GrpcUrl    string
	Checkpoint int64
	logger     log.Logger
}

func NewTwinsRpcServer(rpcUrl string, logger log.Logger) *TwinsRpcServer {
	return &TwinsRpcServer{
		GrpcUrl: rpcUrl,
		logger: logger,
	}
}

func (s *TwinsRpcServer) Heartbeat(ctx context.Context, req *proto.RequestHeartbeat) (*proto.ResponseHeartbeat, error) {
	s.logger.Debug("heartbeat received, timestamp:", req.Timestamp)
	s.Checkpoint = req.Timestamp
	return &proto.ResponseHeartbeat{
		IsElderNow: AsElder,
	}, nil
}

func (s *TwinsRpcServer) Finish(ctx context.Context, req *proto.Request) (*proto.Empty, error) {
	s.logger.Debug("finish received:", req)
	return nil, nil
}

func (s *TwinsRpcServer) Run() error {
	listener, err := net.Listen("tcp", s.GrpcUrl)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterTwinsServer(grpcServer, s)
	return grpcServer.Serve(listener)
}
