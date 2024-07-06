package migrations

import (
	dataarticles "NEWSAPP/features/Articles/dataArticles"
	datacomments "NEWSAPP/features/Comments/dataComments"
	datausers "NEWSAPP/features/Users/dataUsers"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&datausers.Users{})
	db.AutoMigrate(&dataarticles.Articles{})
	db.AutoMigrate(&datacomments.Comments{})
}
