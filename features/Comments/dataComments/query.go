package datacomments

import (
	comments "NEWSAPP/features/Comments"

	"gorm.io/gorm"
)

type commentQuery struct {
	db *gorm.DB
}

// GetAllComments implements comments.DataCommentInterface.


func New(db *gorm.DB) comments.DataCommentInterface {
	return &commentQuery{
		db: db,
	}
}

// CreateComment implements comments.DataCommentInterface.
func (cq *commentQuery) CreateComment(comment comments.Comment) error {
	commentGorm := Comments{
		UserID:     comment.UserID,
		ArticlesID: comment.ArticlesID,
		Content:    comment.Content,
	}
	tx := cq.db.Create(&commentGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetCommentByID implements comments.DataCommentInterface.
func (cq *commentQuery) GetCommentByID(commentID uint) (*comments.Comment, error) {
	var commentData Comments
	tx := cq.db.First(&commentData, commentID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping
	comment := comments.Comment{
		CommentID:  commentData.ID,
		ArticlesID: commentData.ArticlesID,
		Content:    commentData.Content,
	}

	return &comment, nil
}

// UpdateComment implements comments.DataCommentInterface.
func (cq *commentQuery) UpdateComment(commentID uint, comment comments.Comment) error {
	var commentGorm Comments
	tx := cq.db.First(&commentGorm, commentID)
	if tx.Error != nil {
		return tx.Error
	}

	commentGorm.Content = comment.Content

	tx = cq.db.Save(&commentGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteComment implements comments.DataCommentInterface.
func (cq *commentQuery) DeleteComment(commentID uint) error {
	tx := cq.db.Delete(&Comments{}, commentID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (cq *commentQuery) GetAllComments() ([]*comments.Comment, error) {
	var commentsData []*Comments
	tx := cq.db.Find(&commentsData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Prepare slice to hold mapped comments
	var commentsList []*comments.Comment

	// Map each comment from database struct to interface struct
	for _, commentData := range commentsData {
		comment := &comments.Comment{
			CommentID:  commentData.ID,
			ArticlesID: commentData.ArticlesID,
			Content:    commentData.Content,
		}
		commentsList = append(commentsList, comment)
	}

	return commentsList, nil
}