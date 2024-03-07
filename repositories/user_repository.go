package repositories

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/data/entities"
)

type UserRepository interface {
	Create(ctx context.Context, db *sql.DB, entity entities.User) (entities.User, error)
	FindByCredentials(ctx context.Context, db *sql.DB, entity entities.User) (entities.User, error)
	UpdatePassword(tx *sql.Tx, entity entities.User) (entities.User, error)
	ConfirmUser(tx *sql.Tx, entity entities.User) (entities.User, error)
}
