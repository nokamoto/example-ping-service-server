package main

import (
	pb "github.com/nokamoto/example-ping-service"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct{}

func (s *service) Send(_ context.Context, _ *pb.Ping) (*pb.Pong, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *service) SendStreamC(_ pb.PingService_SendStreamCServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *service) SendStreamS(_ *pb.Ping, _ pb.PingService_SendStreamSServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}

func (s *service) SendStreamB(_ pb.PingService_SendStreamBServer) error {
	return status.Error(codes.Unimplemented, "not implemented yet")
}
