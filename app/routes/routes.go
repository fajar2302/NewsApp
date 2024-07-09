package routes

import (
	"NEWSAPP/app/middlewares"
	_articlesData "NEWSAPP/features/Articles/dataArticles"
	_articlesHandler "NEWSAPP/features/Articles/handler"
	_articlesService "NEWSAPP/features/Articles/service"
	_commentData "NEWSAPP/features/Comments/dataComments"
	_commentHandler "NEWSAPP/features/Comments/handler"
	_commentservice "NEWSAPP/features/Comments/service"
	_userData "NEWSAPP/features/Users/dataUsers"
	_userHandler "NEWSAPP/features/Users/handler"
	_userService "NEWSAPP/features/Users/service"

	"NEWSAPP/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	middlewares := middlewares.NewMiddlewares()
	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService, middlewares)
	userHandlerAPI := _userHandler.New(userService)
	commentData := _commentData.New(db)
	commentService := _commentservice.New(commentData)
	commentHandlerAPI := _commentHandler.NewCommentHandler(commentService)
	articlesData := _articlesData.New(db)
	articleService := _articlesService.New(articlesData)
	articlesHandlerAPI := _articlesHandler.New(articleService)

	//userHandler
	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())

	// commentHandler
	e.POST("/comments", commentHandlerAPI.CreateComment, middlewares.JWTMiddleware())
	e.GET("/comments", commentHandlerAPI.ShowAllComments)
	e.DELETE("/comments/:id", commentHandlerAPI.DeleteComment, middlewares.JWTMiddleware())

	//articlesHandler
	e.POST("/articles", articlesHandlerAPI.CreateArtikel, middlewares.JWTMiddleware())
	e.GET("/articles", articlesHandlerAPI.GetAllArtikel)
	e.PUT("/articles/:id", articlesHandlerAPI.UpdateArtikel, middlewares.JWTMiddleware())
	e.DELETE("/articles/:id", articlesHandlerAPI.DeleteArtikel, middlewares.JWTMiddleware())
}
