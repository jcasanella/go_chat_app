package usecase

import (
	"context"
	"time"

	model "github.com/jcasanella/chat_app/model/user"
	repository "github.com/jcasanella/chat_app/repository/user"
)

type UserHandler interface {
	GetUser(ctx context.Context, username string, password string) (model.User, error)
}

type UserService struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewUserService(u repository.UserRepository, t time.Duration) UserHandler {
	return &UserService{
		userRepository: u,
		contextTimeout: t,
	}
}

func (uc *UserService) GetUser(c context.Context, username string, password string) (res model.User, err error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	res, err = uc.userRepository.GetUser(ctx, username, password)
	if err != nil {
		return
	}

	return
}
