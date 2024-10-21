package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

type PGX struct {
	Client *pgxpool.Pool
}

func New(databaseSource string) (store PGX, err error) {
	driverName := strings.ToLower(strings.Split(databaseSource, "://")[0])
	if driverName != "postgres" {
		err = fmt.Errorf("unsupported database driver: %s", driverName)
		return
	}

	store.Client, err = pgxpool.New(context.Background(), databaseSource)
	if err != nil {
		err = fmt.Errorf("error connecting to database: %s", err)
		return
	}

	store.Client.Config().MaxConns = 20

	return
}
