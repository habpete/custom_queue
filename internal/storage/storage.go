package storage

import (
	"context"
	"database/sql"
	"time"

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
	selectStatusQuery = "SELECT id FROM public.statuses WHERE title = $1 LIMIT 1"
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

type UpdateStatusParams struct {
	MessageID int
	Status    string
}

const updateStatusQuery = "UPDATE public.events SET status_id=(SELECT id FROM public.statuses WHERE title = $1 LIMIT 1) WHERE id = $2"

func (i *Storage) UpdateStatus(ctx context.Context, params *UpdateStatusParams) error {
	if _, err := i.db.ExecContext(ctx, updateStatusQuery, params.Status, params.MessageID); err != nil {
		return errors.Wrapf(err, "update status query failed: %v", err)
	}

	return nil
}

type GetMessageParams struct {
	Topic    string
	Statuses []string
	Limit    int
}

type GetMessageResult struct {
	Id        int                        `db:"id"`
	CreatedAt time.Time                  `db:"created_at"`
	Message   map[string]*structpb.Value `db:"message_data"`
}

const getMessageQuery = "SELECT id, created_at, message_data FROM public.events WHERE topic_id = (SELECT id FROM public.topic WHERE title = $1) AND status_id IN (SELECT id FROM public.statuses WHERE title = $2 OR title = $3) LIMIT $4"

func (i *Storage) GetMessage(ctx context.Context, params *GetMessageParams) ([]*GetMessageResult, error) {
	result := make([]*GetMessageResult, params.Limit)
	if err := i.db.SelectContext(ctx, result, getMessageQuery, params.Topic, params.Statuses[0], params.Statuses[1], params.Limit); err != nil {
		return nil, errors.Wrapf(err, "get message query failed: %v", err)
	}

	return result, nil
}
