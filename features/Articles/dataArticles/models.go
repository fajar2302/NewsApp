package dataarticles

import (
	datacomments "NEWSAPP/features/Comments/dataComments"

	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	UserID         uint
	ArtikelPicture string
	ArtikelName    string
	Tag            string
	Description    string
	Comments       []datacomments.Comments `gorm:"foreignKey:ArticlesID"`
}
