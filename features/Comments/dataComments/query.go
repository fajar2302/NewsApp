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

// DeleteComment implements comments.DataCommentInterface.
func (cq *commentQuery) DeleteComment(commentID uint) error {
	tx := cq.db.Delete(&Comments{}, commentID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (cq *commentQuery) GetAllComments() ([]comments.Comment, error) {
	var allComments []Comments // var penampung data yg dibaca dari db
	tx := cq.db.Find(&allComments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allCommentsCore []comments.Comment
	for _, v := range allComments {
		allCommentsCore = append(allCommentsCore, comments.Comment{
			UserID:     v.UserID,
			ArticlesID: v.ArticlesID,
			Content:    v.Content,
		})
	}

	return allCommentsCore, nil
}
