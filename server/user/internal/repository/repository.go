package repository

import (
	"context"
	"qezde/user/internal/config"
	"qezde/user/internal/domain/user"
	"qezde/user/internal/repository/postgres"
)

type Dependencies struct {
	Configs config.Config
	DB      sqlc.DBTX
}

type Configuration func(r *Repository) error

type Repository struct {
	queries *sqlc.Queries
	User    user.Repository
}

func New(d Dependencies, configs ...Configuration) (r *Repository, err error) {
	r = &Repository{
		queries: sqlc.New(d.DB),
	}

	for _, cfg := range configs {
		if err = cfg(r); err != nil {
			return
		}
	}

	return
}

func WithPostgresStore() Configuration {
	return func(r *Repository) (err error) {
		r.User = postgres.NewUserRepository(r.queries, context.Background())

		return
	}
}
