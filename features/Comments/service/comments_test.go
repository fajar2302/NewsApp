package service_test

import (
	comments "NEWSAPP/features/Comments"
	"NEWSAPP/features/Comments/service"
	"NEWSAPP/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateComments(t *testing.T) {
	mockCommentData := new(mocks.DataCommentInterface)
	commentService := service.New(mockCommentData)

	t.Run("success", func(t *testing.T) {
		comment := comments.Comment{
			ArticlesID: 1,
			Content:    "Test comment",
		}

		mockCommentData.On("CreateComment", comment).Return(nil).Once()

		err := commentService.CreateNewComment(comment.ArticlesID, comment)
		assert.NoError(t, err)
		mockCommentData.AssertExpectations(t)
	})

	t.Run("error validation", func(t *testing.T) {
		comment := comments.Comment{} // Empty comment intentionally

		err := commentService.CreateNewComment(0, comment)
		assert.Error(t, err)
		assert.Equal(t, "[validation] ArticlesID, and Content cannot be empty", err.Error())
		mockCommentData.AssertNotCalled(t, "CreateComment") // Ensure CreateComment was not called
	})

	t.Run("error create", func(t *testing.T) {
		comment := comments.Comment{
			ArticlesID: 1,
			Content:    "Test comment",
		}

		mockCommentData.On("CreateComment", comment).Return(errors.New("create error")).Once()

		err := commentService.CreateNewComment(comment.ArticlesID, comment)
		assert.Error(t, err)
		assert.Equal(t, "create error", err.Error())
		mockCommentData.AssertExpectations(t)
	})
}

func TestDeleteComments(t *testing.T) {
	mockCommentData := new(mocks.DataCommentInterface)
	commentService := service.New(mockCommentData)

	t.Run("success", func(t *testing.T) {
		commentID := uint(1)

		mockCommentData.On("DeleteComment", commentID).Return(nil).Once()

		err := commentService.DeleteComment(commentID)
		assert.NoError(t, err)
		mockCommentData.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		commentID := uint(1)

		mockCommentData.On("DeleteComment", commentID).Return(errors.New("delete error")).Once()

		err := commentService.DeleteComment(commentID)
		assert.Error(t, err)
		assert.Equal(t, "delete error", err.Error())
		mockCommentData.AssertExpectations(t)
	})
}

func TestGetAllComments(t *testing.T) {
	mockCommentData := new(mocks.DataCommentInterface)
	commentService := service.New(mockCommentData)

	t.Run("success", func(t *testing.T) {
		expectedComments := []comments.Comment{
			{UserID: 1, ArticlesID: 1, Content: "Comment 1"},
			{UserID: 2, ArticlesID: 1, Content: "Comment 2"},
		}

		mockCommentData.On("GetAllComments").Return(expectedComments, nil).Once()

		commentsList, err := commentService.GetAllComments()
		assert.NoError(t, err)
		assert.Equal(t, expectedComments, commentsList)
		mockCommentData.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockCommentData.On("GetAllComments").Return(nil, errors.New("get error")).Once()

		commentsList, err := commentService.GetAllComments()
		assert.Error(t, err)
		assert.Nil(t, commentsList)
		assert.Equal(t, "get error", err.Error())
		mockCommentData.AssertExpectations(t)
	})
}
