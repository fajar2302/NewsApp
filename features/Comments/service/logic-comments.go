package service

import (
	comments "NEWSAPP/features/Comments"
	"errors"
)

type ServiceComment struct {
	commentData comments.DataCommentInterface
}

func New(cd comments.DataCommentInterface) comments.ServiceCommentInterface {
	return &ServiceComment{
		commentData: cd,
	}
}

func (s *ServiceComment) CreateNewComment(articlesid uint, comment comments.Comment) error {
	// Validate required fields
	if comment.ArticlesID == 0 || comment.Content == "" {
		return errors.New("[validation] ArticlesID, and Content cannot be empty")
	}

	err := s.commentData.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceComment) DeleteComment(commentID uint) error {
	return s.commentData.DeleteComment(commentID)
}
func (s *ServiceComment) GetAllComments() ([]comments.Comment, error) {
	commentsList, err := s.commentData.GetAllComments()
	if err != nil {
		return nil, err
	}
	return commentsList, nil
}
