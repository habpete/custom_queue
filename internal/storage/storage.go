package storage

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
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

type SendMessageParams struct {
	topic   string
	status  string
	message map[string]*structpb.Value
}

const (
	selectTopicQuery  = "SELECT id FROM public.topic WHERE title = $1 LIMIT 1"
	selectStatusQuery = "SELECT status FROM public.statuses WHERE title = $1 LIMIT 1"
	insertEventQuery  = "INSERT INTO public.events (status_id, topic_id, created_at, message_data) VALUES ($1, $2, current_timestamp, $3)"
)

func (i *Storage) SendMessage(ctx context.Context, params *SendMessageParams) error {
	tx, err := i.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly:  false,
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return errors.Wrapf(err, "send message begin tx failed: %v", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	var topicID int
	if err = i.db.GetContext(ctx, &topicID, selectTopicQuery, params.topic); err != nil {
		return errors.Wrapf(err, "select topic query failed: %v", err)
	}

	var statusID int
	if err = i.db.GetContext(ctx, &statusID, selectStatusQuery, params.status); err != nil {
		return errors.Wrapf(err, "select status query failed: %v", err)
	}

	if _, err = i.db.ExecContext(ctx, insertEventQuery, statusID, topicID, params.message); err != nil {
		return errors.Wrapf(err, "insert event query failed: %v", err)
	}

	return nil
}
