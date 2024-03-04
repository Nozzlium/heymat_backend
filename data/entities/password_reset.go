package entities

import "time"

type PasswordReset struct {
	Ticket    string
	UserID    uint32
	Redeemed  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpireAt  time.Time
}
