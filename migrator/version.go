package main

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	selectTableVersionIsExist = "SELECT 1 FROM information_schema.tables WHERE scheme = 'public' AND name = 'db_version'"
	createTableVersion        = "CREATE TABLE public.db_version (id SERIAL, name TEXT, created_at TIMESTAMP, processing_at TIMESTAMP, status TEXT)"
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

const createStatusQuery = ""

func createStatus(ctx context.Context, conn *sqlx.DB, name, status string) error {
	if conn == nil {
		return fmt.Errorf("")
	}

	_, err := conn.ExecContext(ctx, createStatusQuery)
	if err != nil {
		return errors.Wrapf(err, "create status failed: %v", err)
	}

	return nil
}

const updateStatusQuery = ""

func updateStatus(ctx context.Context, conn *sqlx.DB, status string) error {
	if conn == nil {
		return fmt.Errorf("")
	}

	_, err := conn.ExecContext(ctx, updateStatusQuery, status)
	if err != nil {
		return errors.Wrapf(err, "update status failed: %v", err)
	}

	return nil
}
