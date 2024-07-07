package handler

type CommentRequest struct {
	// UserID     uint   `json:"user_id"`
	ArticlesID uint   `json:"articles_id"`
	Content    string `json:"content"`
}
type UpdateRequest struct {
	Content    string `json:"content"`
}
