package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConnectParams struct {
	host     string `json:"host"`
	port     int    `json:"port"`
	database string `json:"database"`
	username string `json:"username"`
	password string `json:"password"`
}

const connectionStringPattern = "postgres://%s:%s@%s:%d/%s"

func buildConnectionString(params *ConnectParams) string {
	return fmt.Sprintf(connectionStringPattern, params.username, params.password, params.host, params.port, params.database)
}

func connect(params *ConnectParams) (*sqlx.DB, error) {
	connString := buildConnectionString(params)
	if connString == "" {
		return nil, fmt.Errorf("empty connection string for postgres")
	}

	conn, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
