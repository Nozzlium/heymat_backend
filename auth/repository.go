package auth

import (
	"context"
	"database/sql"
	"log"
)

func save(
	ctx context.Context,
	DB *sql.DB,
	userAccount UserAccount,
) (UserAccount, error) {
	query := `
    insert 
    into user_account (
      username, email, password, is_email_confirmed, created_at, updated_at
    ) 
    values (
      $1, $2, $3, $4, $5, $6
    ) returning id; 
  `
	var insertedId uint64
	err := DB.QueryRowContext(
		ctx,
		query,
		userAccount.Username,
		userAccount.Email,
		userAccount.Password,
		userAccount.IsEmailConfirmed,
		userAccount.CreatedAt,
		userAccount.UpdatedAt,
	).Scan(&insertedId)
	if err != nil {
		log.Println(err)
		return UserAccount{}, err
	}

	userAccount.ID = insertedId
	return userAccount, nil
}

func findById(
	ctx context.Context,
	DB *sql.DB,
	userId uint64,
) (UserAccount, error) {
	query := `
    select 
      id,
      username,
      email
    from user_account
    where id = $1;
  `
	user := UserAccount{}
	err := DB.QueryRowContext(ctx, query, userId).
		Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Println(err)
		return UserAccount{}, err
	}

	return user, nil
}