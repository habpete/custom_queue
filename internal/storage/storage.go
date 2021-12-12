package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func (i *Storage) SendMessage() error {
	return nil
}
