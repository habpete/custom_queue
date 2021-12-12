package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Storage struct {
	db *sqlx.DB
}

func New(params *ConnectParams) (*Storage, error) {
	conn, err := connect(params)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: conn,
	}, nil
}

type SendMessageParams struct{
	topic string
	status string
}

const sendMessageQuery = "
INSERT INTO public.events (created_at, message_data) VALUES (current_timestamp, $1)
"

func (i *Storage) SendMessage(ctx context.Context, params *SendMessageParams) error {
	_, err := i.db.ExecContext(ctx, sendMessageQuery)
	if err != nil {
		return errors.Wrapf(err, "send message query failed: %v", err)
	}

	return nil
}
