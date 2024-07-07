package comments

type Comment struct {
	CommentID  uint   `json:"comment_id"`
	UserID     uint   `json:"user_id"`
	ArticlesID uint   `json:"article_id"`
	Content    string `json:"content"`
}

type DataCommentInterface interface {
	CreateComment(comment Comment) error
	GetCommentByID(commentID uint) (*Comment, error)
	UpdateComment(commentID uint, newComment Comment) error
	DeleteComment(commentID uint) error
	GetAllComments() ([]*Comment, error)
}

type ServiceCommentInterface interface {
	CreateNewComment(userId uint, comment Comment) error
	GetCommentDetails(commentID uint) (*Comment, error)
	UpdateCommentDetails(commentID uint, updatedComment Comment) error
	DeleteComment(commentID uint) error
	GetAllComments() ([]*Comment, error)
}