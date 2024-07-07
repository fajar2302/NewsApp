package handler

import (
	"net/http"
	"strconv"

	"NEWSAPP/app/middlewares"
	comments "NEWSAPP/features/Comments"
	"NEWSAPP/utils/responses"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentService comments.ServiceCommentInterface
}

func NewCommentHandler(cs comments.ServiceCommentInterface) *CommentHandler {
	return &CommentHandler{
		commentService: cs,
	}
}

func (ch *CommentHandler) CreateComment(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "error", "Unauthorized", nil))
	}

	// Bind request body to struct
	newComment := CommentRequest{}
	if err := c.Bind(&newComment); err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Error binding data: "+err.Error(), nil))
	}
	
	// Mapping request ke struct Comment
	dataComments := comments.Comment{
		ArticlesID:    newComment.ArticlesID,
		Content:       newComment.Content,
	}

	if err := ch.commentService.CreateNewComment(uint(userID), dataComments); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "Comment creation failed: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "Success", "Comment created successfully", nil))
}

func (ch *CommentHandler) ShowAllComments(c echo.Context) error {
	commentsList, err := ch.commentService.GetAllComments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "Failed to fetch comments: "+err.Error(), nil))
	}

	// Prepare slice to hold mapped comments with reduced fields
	var commentsResponse []CommentResponse

	for _, comment := range commentsList {
		commentResponse := CommentResponse{
			ID:      comment.CommentID,
			ArticlesID: comment.ArticlesID,
			Content: comment.Content,
		}
		commentsResponse = append(commentsResponse, commentResponse)
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "Success", "All comments fetched successfully", commentsResponse))
}

func (ch *CommentHandler) ShowCommentByID(c echo.Context) error {
	// Mengambil commentID dari URL parameter dan mengonversi ke uint
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "Bad Request", "Invalid comment ID", nil))
	}

	comment, err := ch.commentService.GetCommentDetails(uint(commentID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "Failed to fetch comment: "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "Success", "Comment fetched successfully", comment))
}



func (ch *CommentHandler) UpdateComment(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "error", "Unauthorized", nil))
	}

	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "Bad Request", "Invalid comment ID", nil))
	}

	updatedComment := UpdateRequest{}
	if err := c.Bind(&updatedComment); err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Error binding data: "+err.Error(), nil))
	}

	if err := ch.commentService.UpdateCommentDetails(uint(commentID), comments.Comment{
		Content: updatedComment.Content,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "Failed to update comment: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "Success", "Comment updated successfully", nil))
}
func (ch *CommentHandler) DeleteComment(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "error", "Unauthorized", nil))
	}

	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "Bad Request", "Invalid comment ID", nil))
	}

	if err := ch.commentService.DeleteComment(uint(commentID)); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "Failed to delete comment: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "Success", "Comment deleted successfully", nil))
}
