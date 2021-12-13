package producer

import (
	"context"

	pb "github.com/custom_queue/pkg/proto"
	"google.golang.org/grpc"
)

type Producer struct {
	conn   *grpc.ClientConn
	client pb.MessageQueueClient
}

func New(grpcHost string) (*Producer, error) {
	conn, err := grpc.Dial(grpcHost, grpc.WithInsecure)
	if err != nil {
		return nil, err
	}

	return &Producer{
		conn:   conn,
		client: pb.NewMessageQueueClient(conn),
	}, nil
}

func (i *Producer) Close() error {
	return i.conn.Close()
}

func (i *Producer) SetDone(ctx context.Context, req *pb.SetDoneRequest) (*pb.SetDoneResponse, error) {
	return i.client.SetDone(ctx, req)
}

func (i *Producer) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	return i.client.GetMessages(ctx, req)
}
