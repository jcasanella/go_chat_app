package model

import (
	"context"
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

type UserUsecase interface {
	GetUser(ctx context.Context, username string, password string) (User, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, username string, password string) (res User, err error)
}
