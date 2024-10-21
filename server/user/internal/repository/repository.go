package repository

import (
	"qezde/user/internal/config"
	"qezde/user/internal/domain/user"
	"qezde/user/internal/repository/postgres"
	"qezde/user/pkg/store"
)

type Dependencies struct {
	Configs config.Config
}

type Configuration func(r *Repository) error

type Repository struct {
	postgres store.PGX

	User user.Repository
}

func New(d Dependencies, configs ...Configuration) (r *Repository, err error) {
	r = &Repository{}

	for _, cfg := range configs {
		if err = cfg(r); err != nil {
			return
		}
	}

	return
}

func WithPostgresStore(databaseSource string) Configuration {
	return func(r *Repository) (err error) {
		r.postgres, err = store.New(databaseSource)
		if err != nil {
			return
		}

		r.User = postgres.NewUserRepository(r.postgres.Client)

		return
	}
}
