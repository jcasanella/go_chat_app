package usecase

import (
	"context"
	"time"

	model "github.com/jcasanella/chat_app/model/user"
)

type UserUsecase struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u model.UserRepository, t time.Duration) model.UserUsecase {
	return &UserUsecase{
		userRepository: u,
		contextTimeout: t,
	}
}

func (uc *UserUsecase) GetUser(c context.Context, username string, password string) (res model.User, err error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	res, err = uc.userRepository.GetUser(ctx, username, password)
	if err != nil {
		return
	}

	return
}
