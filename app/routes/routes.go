package routes

import (
	"NEWSAPP/app/middlewares"
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
	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)
	commentData := _commentData.New(db)
	commentService := _commentservice.New(commentData)
	commentHandlerAPI := _commentHandler.NewCommentHandler(commentService)

	//userHandler
	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())

	// commentHandler
	e.POST("/commets", commentHandlerAPI.CreateComment, middlewares.JWTMiddleware())
	e.GET("/commets", commentHandlerAPI.ShowAllComments, middlewares.JWTMiddleware())
	e.PUT("/commets/:id", commentHandlerAPI.UpdateComment, middlewares.JWTMiddleware())
	e.DELETE("/commets/:id", commentHandlerAPI.DeleteComment, middlewares.JWTMiddleware())
}
