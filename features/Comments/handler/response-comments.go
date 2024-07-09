package handler

import model "NEWSAPP/features/Comments/dataComments"

type CommentResponse struct {
	UserID     uint   `json:"user_id"`
	ArticlesID uint   `json:"articles_id"`
	Content    string `json:"content"`
}

func NewCommentResponse(comments []*model.Comments) []CommentResponse {
	var responses []CommentResponse

	for _, comment := range comments {
		response := CommentResponse{
			UserID:     comment.UserID,
			ArticlesID: comment.ArticlesID,
			Content:    comment.Content,
		}
		responses = append(responses, response)
	}

	return responses
}
