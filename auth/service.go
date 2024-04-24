package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/nozzlium/heymat_backend/lib"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

var (
	ErrUserExists = errors.New(
		"username atau email sudah terdaftar",
	)
	ErrInvalidCredential = errors.New(
		"username, email, atau password salah",
	)
)

var loginErrors []error = []error{
	sql.ErrNoRows,
	bcrypt.ErrMismatchedHashAndPassword,
}

type UserResponse struct {
	ID               uint64 `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	IsEmailConfirmed bool   `json:"isEmailConfirmed"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func registerUser(
	ctx context.Context,
	user UserAccount,
) (UserResponse, error) {
	passwordHash, err := hashPassword(
		user.Password,
	)
	if err != nil {
		log.Println(err)
		return UserResponse{}, err
	}

	currentTime := time.Now()

	user.Password = passwordHash
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	result, err := save(
		ctx,
		DB,
		user,
	)
	if err != nil {
		return UserResponse{}, parseRegisterError(
			err,
		)
	}

	return result.toResponse(), nil
}

func login(
	ctx context.Context,
	userAccount UserAccount,
) (LoginResponse, error) {
	user, err := findUserByUsernameOrEmail(
		ctx,
		DB,
		userAccount,
	)
	if err != nil {
		return LoginResponse{}, parseLoginError(
			err,
		)
	}

	err = compareHashWithPassword(
		user.Password,
		userAccount.Password,
	)
	if err != nil {
		log.Println(err)
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return LoginResponse{}, ErrInvalidCredential
		}
		return LoginResponse{}, err
	}

	token, err := lib.GenerateJwtToken(
		lib.InitAuthClaims(
			user.ID,
			user.Username,
			user.Email,
		),
	)
	if err != nil {
		log.Println(err)
		return LoginResponse{}, err
	}

	return LoginResponse{
		Token: token,
	}, nil
}

func parseRegisterError(
	err error,
) error {
	if driverErr, ok := err.(*pq.Error); ok {
		switch driverErr.Code.Name() {
		case "unique_violation":
			return ErrUserExists
		}
	}
	return err
}

func parseLoginError(err error) error {
	for _, correspErr := range loginErrors {
		if err == correspErr {
			return ErrInvalidCredential
		}
	}
	return err
}
