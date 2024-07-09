package handler

type CommentRequest struct {
	ArticlesID uint   `json:"articles_id" form:"articles_id"`
	Content    string `json:"content" form:"content"`
}
