package mocks

import (
	"context"

	model "github.com/jcasanella/chat_app/model/user"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (_m *UserRepository) GetUser(ctx context.Context, username string, password string) (res model.User, err error) {
	ret := _m.Called(ctx, username, password)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) model.User); ok {
		r0 = rf(ctx, username, password)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
