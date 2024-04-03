package auth

import (
	"context"
	"database/sql"
)

func getUserByCredential(
	ctx context.Context,
	DB *sql.DB,
	entity UserAccount,
) (UserAccount, error) {
	query := `
    select
      id,
      username,
      email,
      is_email_confirmed,
      password
    from user_account 
    where username = $1 or email = $2
    limit 1
  `
	user := UserAccount{}
	err := DB.QueryRowContext(ctx, query, entity.Username, entity.Email).
		Scan(&user.ID, &user.Username, &user.Email, &user.IsEmailConfirmed, &user.Password)
	return user, err
}
