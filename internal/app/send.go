package app

import (
	"context"

	"github.com/custom_queue/internal/storage"
	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Service) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "incorrect request parameters")
	}

	if err := i.strg.SendMessage(ctx, &storage.SendMessageParams{}); err != nil {
		return nil, status.Error(codes.Internal, "send failed")
	}

	return &pb.SendResponse{}, nil
}
