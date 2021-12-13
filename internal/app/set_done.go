package app

import (
	"context"

	"github.com/custom_queue/internal/storage"
	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Service) SetDone(ctx context.Context, req *pb.SetDoneRequest) (*pb.SetDoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "incorrect request parameters")
	}

	if err := i.strg.UpdateStatus(ctx, &storage.UpdateStatusParams{
		MessageID: int(req.MessageId),
		Status:    "Success",
	}); err != nil {
		return nil, status.Error(codes.Internal, "set done update status failed")
	}

	return &pb.SetDoneResponse{}, nil
}
