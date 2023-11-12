package usecase

import (
	"context"
	"time"

	model "github.com/jcasanella/chat_app/model/user"
)

type userUsecase struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u model.UserRepository, t time.Duration) model.UserUsecase {
	return &userUsecase{
		userRepository: u,
		contextTimeout: t,
	}
}

func (uc *userUsecase) GetUser(c context.Context, username string, password string) (res model.User, err error) {
	// ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	// defer cancel()

	res, err = uc.userRepository.GetUser(c, username, password)
	if err != nil {
		return
	}

	return
}
