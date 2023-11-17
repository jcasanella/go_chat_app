package usecase

import (
	"context"
	"testing"
	"time"

	model "github.com/jcasanella/chat_app/model/user"
	mocks "github.com/jcasanella/chat_app/model/user/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	mockRepos := mocks.UserRepository{}

	expectedUser := model.User{
		ID:        "MockId",
		Name:      "test",
		Username:  "test",
		Password:  "test",
		CreatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	mockRepos.On("GetUser", mock.Anything, "test", "test").Return(expectedUser, nil)

	userService := NewUserService(&mockRepos, 5*time.Second)
	actualUser, err := userService.GetUser(context.TODO(), "test", "test")

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, expectedUser, actualUser, "User should be valid")
}
