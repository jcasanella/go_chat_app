package repository

import (
	"context"
	"testing"
	"time"

	model "github.com/jcasanella/chat_app/model/user"
	mocks "github.com/jcasanella/chat_app/model/user/mock"
	usecase "github.com/jcasanella/chat_app/usecase/user"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	mockRepos := mocks.UserRepository{}

	userExp := model.User{
		ID:        "MockId",
		Name:      "test",
		Username:  "test",
		Password:  "test",
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	mockRepos.On("GetUser", context.TODO(), "test", "test").Return(userExp, nil)

	userUsecase := usecase.NewUserUsecase(&mockRepos, 5*time.Second)
	userAct, err := userUsecase.GetUser(context.TODO(), "test", "test")

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, userExp, userAct, "User should be valid")
}
