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

func (s *ServiceComment) CreateNewComment(userID uint, comment comments.Comment) error {
	// Validate required fields
	if userID == 0 || comment.ArticlesID == 0 || comment.Content == "" {
		return errors.New("[validation] UserID, ArticlesID, dan Content tidak boleh kosong")
	}

	comment.UserID = userID

	return s.commentData.CreateComment(comment)
}
func (s *ServiceComment) GetCommentDetails(commentID uint) (*comments.Comment, error) {
	comment, err := s.commentData.GetCommentByID(commentID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (s *ServiceComment) UpdateCommentDetails(commentID uint, updatedComment comments.Comment) error {
	if updatedComment.Content == "" {
		return errors.New("[validation] Content tidak boleh kosong")
	}

	return s.commentData.UpdateComment(commentID, updatedComment)
}

func (s *ServiceComment) DeleteComment(commentID uint) error {
	return s.commentData.DeleteComment(commentID)
}
func (s *ServiceComment) GetAllComments() ([]*comments.Comment, error) {
	commentsList, err := s.commentData.GetAllComments()
	if err != nil {
		return nil, err
	}
	return commentsList, nil
}
