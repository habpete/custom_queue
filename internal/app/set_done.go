package app

import (
	"context"

	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Service) SetDone(ctx context.Context, req *pb.SetDoneRequest) (*pb.SetDoneResponse, error) {
	return nil, status.Error(codes.Unimplemented, "set done is not implemented")
}
