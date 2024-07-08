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
			UserID:     1,
			ArticlesID: 1,
			Content:    "Great article!",
		}

		mockCommentData.On("CreateComment", comment).Return(nil).Once()

		err := commentService.CreateNewComment(1, comment)
		assert.NoError(t, err)
		mockCommentData.AssertExpectations(t)
	})

	t.Run("error validation", func(t *testing.T) {
		comment := comments.Comment{
			ArticlesID: 1,
			Content:    "Great article!",
		}

		err := commentService.CreateNewComment(0, comment)
		assert.Error(t, err)
		assert.Equal(t, "[validation] UserID, ArticlesID, dan Content tidak boleh kosong", err.Error())
		mockCommentData.AssertExpectations(t)
	})

	t.Run("error create", func(t *testing.T) {
		comment := comments.Comment{
			UserID:     1,
			ArticlesID: 1,
			Content:    "Great article!",
		}

		mockCommentData.On("CreateComment", comment).Return(errors.New("create error")).Once()

		err := commentService.CreateNewComment(1, comment)
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
		expectedComments := []*comments.Comment{
			{
				CommentID:  1,
				ArticlesID: 1,
				Content:    "Great article!",
			},
			{
				CommentID:  2,
				ArticlesID: 2,
				Content:    "Nice post!",
			},
		}

		mockCommentData.On("GetAllComments").Return(expectedComments, nil).Once()

		returnedComments, err := commentService.GetAllComments()
		assert.NoError(t, err)
		assert.Equal(t, expectedComments, returnedComments)
		mockCommentData.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockCommentData.On("GetAllComments").Return(nil, errors.New("get all error")).Once()

		returnedComments, err := commentService.GetAllComments()
		assert.Error(t, err)
		assert.Equal(t, "get all error", err.Error())
		assert.Nil(t, returnedComments)
		mockCommentData.AssertExpectations(t)
	})
}
