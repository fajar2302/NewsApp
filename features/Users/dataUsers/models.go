package datausers

import (
	dataarticles "NEWSAPP/features/Articles/dataArticles"
	datacomments "NEWSAPP/features/Comments/dataComments"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ProfilePicture string
	FullName       string
	Email          string `gorm:"unique"`
	Password       string
	PhoneNumber    string
	Address        string
	Articles       []dataarticles.Articles `gorm:"foreignKey:UserID"`
	Comments       []datacomments.Comments `gorm:"foreignKey:UserID"`
}
