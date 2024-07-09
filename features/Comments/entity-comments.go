package comments

type Comment struct {
	CommentID  uint
	UserID     uint
	ArticlesID uint
	Content    string
}

type DataCommentInterface interface {
	CreateComment(comment Comment) error
	DeleteComment(commentID uint) error
	GetAllComments() ([]Comment, error)
}

type ServiceCommentInterface interface {
	CreateNewComment(articlesid uint, comment Comment) error
	DeleteComment(commentID uint) error
	GetAllComments() ([]Comment, error)
}
