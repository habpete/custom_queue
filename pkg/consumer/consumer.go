package consumer

import (
	"context"

	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc"
)

type Consumer struct {
	conn   *grpc.ClientConn
	client pb.MessageQueueClient
}

func New(grpcHost string) (*Consumer, error) {
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:   conn,
		client: pb.NewMessageQueueClient(conn),
	}, nil
}

func (i *Consumer) Close() error {
	return i.conn.Close()
}

func (i *Consumer) Send(ctx context.Context, req *pb.SendRequest) error {
	_, err := i.client.Send(ctx, req)
	return err
}
