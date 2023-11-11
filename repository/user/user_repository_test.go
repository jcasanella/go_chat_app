package repository

import (
	"context"
	"testing"
	"time"

	model "github.com/jcasanella/chat_app/model/user"
	mocks "github.com/jcasanella/chat_app/model/user/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	mockRepos := mocks.UserRepository{}
	mockRepos.On("GetUser", "test", "test").Return(&model.User{
		ID:        "MockId",
		Name:      "test",
		Password:  "test",
		CreatedAt: nil,
		UpdatedAt: nil,
	}, nil)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "updated_at", "created_at"}).
		AddRow(1, "title 1", "Content 1", 1, time.Now(), time.Now())

	query := "SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := articleMysqlRepo.NewMysqlArticleRepository(db)

	num := int64(5)
	anArticle, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
