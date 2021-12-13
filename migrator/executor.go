package main

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func migrate(ctx context.Context, pathToMigrations string) error {
	files, err := os.ReadDir(pathToMigrations)
	if err != nil {
		return errors.Wrapf(err, "read directory failed: %v", err)
	}

	for _, file := range files {
		if err = execute(ctx, path.Join(pathToMigrations, file.Name())); err != nil {
			return errors.Wrapf(err, "migration for file %s failed: %v", file.Name(), err)
		}
	}

	return nil
}

func execute(ctx context.Context, pathToFile string) error {
	if pathToFile == "" {
		return fmt.Errorf("empty path to sql file")
	}

	conn, err := sqlx.Connect("postgres", "")
	if err != nil {
		return errors.Wrapf(err, "connect to database was failed: %v", err)
	}

	defer conn.Close()

	if _, err = conn.ExecContext(ctx, ""); err != nil {
		return errors.Wrapf(err, "exec migration script was failed: %v", err)
	}

	return nil
}
