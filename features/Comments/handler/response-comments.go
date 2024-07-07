package handler

import model "NEWSAPP/features/Comments/dataComments"

type CommentResponse struct {
	ID         uint   `json:"id"`
	ArticlesID uint   `json:"articles_id"`
	Content    string `json:"content"`
}

func NewCommentResponse(comments []*model.Comments) []CommentResponse {
	var responses []CommentResponse

	for _, comment := range comments {
		response := CommentResponse{
			ID:         comment.ID,
			ArticlesID: comment.ArticlesID,
			Content:    comment.Content,
		}
		responses = append(responses, response)
	}

	return responses
}
