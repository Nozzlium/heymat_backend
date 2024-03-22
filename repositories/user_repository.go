package repositories

import (
	"context"
	"database/sql"

	"github.com/nozzlium/heymat_backend/entities"
)

type UserRepository interface {
	Create(
		ctx context.Context,
		db *sql.DB,
		entity entities.UserAccount,
	) (entities.UserAccount, error)
	FindByCredentials(
		ctx context.Context,
		db *sql.DB,
		entity entities.UserAccount,
	) (entities.UserAccount, error)
	UpdatePassword(tx *sql.Tx, entity entities.UserAccount) (entities.UserAccount, error)
	ConfirmUser(tx *sql.Tx, entity entities.UserAccount) (entities.UserAccount, error)
}
