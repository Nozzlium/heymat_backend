package entities

import "time"

type User struct {
	ID               uint64
	Username         string
	Email            string
	Password         string
	IsEmailConfirmed bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
