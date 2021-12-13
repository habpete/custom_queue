package app

import (
	"context"

	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Service) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "get messages is not implemented")
}
