package entities

import "time"

type User struct {
	ID               uint32
	Username         string
	Email            string
	Password         string
	IsEmailConfirmed bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
