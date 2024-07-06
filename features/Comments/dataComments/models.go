package datacomments

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	UserID     uint
	ArticlesID uint
	Content    string
}
