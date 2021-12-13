package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	selectTableVersionIsExist = "SELECT 1 FROM information_schema.tables WHERE scheme = 'public' AND name = 'db_version'"
	createTableVersion        = ""
)

func checkVersion(ctx context.Context, connectionString string) error {
	conn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return errors.Wrapf(err, "connect to database failed: %v", err)
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, selectTableVersionIsExist)
	if err != nil {
		return errors.Wrapf(err, "check table version is exist failed: %v", err)
	}

	if rows.Next() {
		return nil
	}

	if _, err = conn.ExecContext(ctx, createTableVersion); err != nil {
		return errors.Wrapf(err, "create table version failed: %v", err)
	}

	return nil
}
