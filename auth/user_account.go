package auth

import "time"

type UserAccount struct {
	ID               uint64
	Username         string
	Email            string
	Password         string
	IsEmailConfirmed bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (userAccount UserAccount) toResponse() UserResponse {
	return UserResponse{
		ID:               userAccount.ID,
		Username:         userAccount.Username,
		Email:            userAccount.Email,
		IsEmailConfirmed: userAccount.IsEmailConfirmed,
	}
}
