package app

import (
	"context"

	"github.com/custom_queue/internal/storage"
	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Service) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	if req == nil || req.Topic == "" {
		return nil, status.Error(codes.InvalidArgument, "incorrect request parameters")
	}

	result, err := i.strg.GetMessage(ctx, &storage.GetMessageParams{})
	if err != nil {
		return nil, status.Error(codes.Internal, "get messages failed")
	}

	if len(result) == 0 {
		return nil, status.Error(codes.NotFound, "get messages is empty")
	}

	return &pb.GetMessagesResponse{
		MessageId: int64(result[0].Id),
		Message:   result[0].Message,
	}, nil
}
