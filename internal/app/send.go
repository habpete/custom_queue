package app

import (
	"context"

	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Service) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	return nil, status.Error(codes.Unimplemented, "send is not implemented")
}
