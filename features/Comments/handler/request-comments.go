package handler

type CommentRequest struct {
	ArticlesID uint   `json:"articles_id"`
	Content    string `json:"content"`
}
