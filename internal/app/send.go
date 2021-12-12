package app

import (
	"context"
	"fmt"

	pb "github.com/custom_queue/pkg/proto"
)

func (Service) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	return nil, fmt.Errorf("send is not implemented")
}
