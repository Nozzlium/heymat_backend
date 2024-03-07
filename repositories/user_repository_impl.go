package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nozzlium/heymat_backend/data/entities"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(
	ctx context.Context,
	db *sql.DB,
	entity entities.User,
) (entities.User, error) {
	query := `
    insert 
    into users (
      username, email, password, is_email_confirmed, created_at
    ) 
    values (
      $1, $2, $3, $4, $5
    ) returning id; 
  `
	var insertedId uint32
	err := db.QueryRowContext(
		ctx,
		query,
		entity.Username,
		entity.Email,
		entity.Password,
		entity.IsEmailConfirmed,
		time.Now(),
	).Scan(&insertedId)
	if err != nil {
		return entities.User{}, err
	}

	entity.ID = uint32(insertedId)
	return entity, nil
}

func (repository *UserRepositoryImpl) FindByCredentials(
	ctx context.Context,
	db *sql.DB,
	entity entities.User,
) (entities.User, error) {
	query := `
    select
      id,
      username,
      email,
      is_email_confirmed,
      password
    from users
    where username = $1 or email = $2
    limit 1
  `
	user := entities.User{}
	row := db.QueryRowContext(ctx, query, entity.Username, entity.Email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.IsEmailConfirmed, &user.Password)
	return user, err
}

func (repository *UserRepositoryImpl) UpdatePassword(
	tx *sql.Tx,
	entity entities.User,
) (entities.User, error) {
	query := `
    alter
    table users
    (password)
    values (?)
  `
	result, err := tx.Exec(query, entity.Password)
	if err != nil {
		return entities.User{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entities.User{}, err
	}

	if rowsAffected < 1 {
		return entities.User{}, errors.New("unknown error")
	}

	return entity, nil
}

func (repository *UserRepositoryImpl) ConfirmUser(
	tx *sql.Tx,
	entity entities.User,
) (entities.User, error) {
	query := `
    alter
    table users
    (is_email_confirmed)
    values (true)
  `
	result, err := tx.Exec(query)
	if err != nil {
		return entities.User{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entities.User{}, err
	}

	if rowsAffected < 1 {
		return entity, errors.New("unknown error, alter unsuccessful")
	}

	return entity, nil
}
