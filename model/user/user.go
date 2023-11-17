package model

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
