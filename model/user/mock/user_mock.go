package mocks

import (
	"context"

	model "github.com/jcasanella/chat_app/model/user"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (_m *UserRepository) GetUser(ctx context.Context, username string, password string) (model.User, error) {
	ret := _m.Called(ctx, username, password)
	return ret.Get(0).(model.User), ret.Error(1)
}
