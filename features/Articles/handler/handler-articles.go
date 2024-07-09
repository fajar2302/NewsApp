package handler

import (
	"NEWSAPP/app/middlewares"
	articles "NEWSAPP/features/Articles"
	"NEWSAPP/utils/responses"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ArtikelHandler struct {
	artikelService articles.ServiceArtikelInterface
}

func New(ah articles.ServiceArtikelInterface) *ArtikelHandler {
	return &ArtikelHandler{
		artikelService: ah,
	}
}

func (ah *ArtikelHandler) CreateArtikel(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	// membaca data dari request body
	newArtikel := ArtikelRequest{}
	errBind := c.Bind(&newArtikel)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error bind artikel: " + errBind.Error(),
		})
	}

	// Membaca file gambar pengguna (jika ada)
	file, err := c.FormFile("articles_picture")
	var imageURL string
	if err == nil {
		// Buka file
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Gagal membuka file gambar: "+err.Error(), nil))
		}
		defer src.Close()

		// Upload file ke Cloudinary
		imageURL, err = newArtikel.uploadToCloudinary(src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Gagal mengunggah gambar: "+err.Error(), nil))
		}
	}

	// mapping  dari request ke articles
	inputArtikel := articles.Artikel{
		UserID:         uint(userID),
		ArtikelPicture: imageURL,
		ArtikelName:    newArtikel.ArtikelName,
		Tag:            newArtikel.Tag,
		Description:    newArtikel.Description,
	}

	if errInsert := ah.artikelService.Create(inputArtikel); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Article News failed to be created: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Article News failed to be created: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "Article News was successfully created", nil))
}

func (ah *ArtikelHandler) GetAllArtikel(c echo.Context) error {
	result, err := ah.artikelService.GetAllArtikel()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "failed",
			"message": "Failed to retrieve all article news",
		})
	}
	var allArtikelResponse []ArtikelResponse
	for _, value := range result {
		allArtikelResponse = append(allArtikelResponse, ArtikelResponse{
			UserID:         value.UserID,
			ArtikelPicture: value.ArtikelPicture,
			ArtikelName:    value.ArtikelName,
			Tag:            value.Tag,
			Description:    value.Description,
		})
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "All articles fetched successfully", allArtikelResponse))
}

func (ah *ArtikelHandler) DeleteArtikel(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}
	if errInsert := ah.artikelService.Delete(uint(idConv), uint(userID)); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Failed to delete the article news: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Failed to delete the article news: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "Successfully deleted a article news", nil))
}

func (ah *ArtikelHandler) UpdateArtikel(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error converting id: " + errConv.Error(),
		})
	}

	var updateData ArtikelRequest
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error binding todo: " + err.Error(),
		})
	}

	// Membaca file gambar pengguna (jika ada)
	file, err := c.FormFile("articles_picture")
	var imageURL string
	if err == nil {
		// Buka file
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Gagal membuka file gambar: "+err.Error(), nil))
		}
		defer src.Close()

		// Upload file ke Cloudinary
		imageURL, err = updateData.uploadToCloudinary(src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Gagal mengunggah gambar: "+err.Error(), nil))
		}
	}

	inputArtikel := articles.Artikel{
		ArtikelPicture: imageURL,
		ArtikelName:    updateData.ArtikelName,
		Tag:            updateData.Tag,
		Description:    updateData.Description,
	}

	if errInsert := ah.artikelService.Update(uint(idConv), uint(userID), inputArtikel); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Failed to update the article news: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Failed to update the article news: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "Successfully updated the article news", nil))
}
