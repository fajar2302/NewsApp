package comments

type Comment struct {
	CommentID  uint   `json:"comment_id"`
	UserID     uint   `json:"user_id"`
	ArticlesID uint   `json:"article_id"`
	Content    string `json:"content"`
}

type DataCommentInterface interface {
	CreateComment(comment Comment) error
	DeleteComment(commentID uint) error
	GetAllComments() ([]*Comment, error)
}

type ServiceCommentInterface interface {
	CreateNewComment(userId uint, comment Comment) error
	DeleteComment(commentID uint) error
	GetAllComments() ([]*Comment, error)
}
