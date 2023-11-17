package repository

import (
	"context"

	"gorm.io/gorm"

	model "github.com/jcasanella/chat_app/model/user"
)

type UserRepository interface {
	GetUser(ctx context.Context, username string, password string) (res model.User, err error)
}

type DBUserRepository struct {
	conn *gorm.DB
}

func NewDBUserRepository(conn *gorm.DB) UserRepository {
	return &DBUserRepository{conn}
}

func (m *DBUserRepository) GetUser(ctx context.Context, username string, password string) (model.User, error) {
	var res model.User
	err := m.conn.WithContext(ctx).Where("username = ? AND password = ?", username, password).First(&res).Error

	return res, err
}
