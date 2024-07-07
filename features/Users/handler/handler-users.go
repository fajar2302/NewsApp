package handler

import (
	"NEWSAPP/app/middlewares"
	users "NEWSAPP/features/Users"
	"NEWSAPP/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.ServiceUserInterface
}

func New(us users.ServiceUserInterface) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	// Membaca data dari body permintaan
	newUser := UserRequest{}
	if errBind := c.Bind(&newUser); errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Error binding data: "+errBind.Error(), nil))
	}

	// Mapping request ke struct User
	dataUser := users.User{
		FullName:    newUser.FullName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		PhoneNumber: newUser.PhoneNumber,
		Address:     newUser.Address,
	}

	// Memanggil service layer untuk menyimpan data
	if errInsert := uh.userService.RegistrasiAccount(dataUser); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "Failed", "User registration failed: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "User registration failed: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "Success", "User registration succcessful", nil))
}

func (uh *UserHandler) Login(c echo.Context) error {
	// Membaca data dari request body
	loginReq := LoginRequest{}

	if errBind := c.Bind(&loginReq); errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Error binding data: "+errBind.Error(), nil))
	}

	// Melakukan login
	_, token, err := uh.userService.LoginAccount(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "User login failed: "+err.Error(), nil))
	}

	// Mengembalikan respons dengan token
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "Success", "User login successful", echo.Map{"token": token}))
}

func (uh *UserHandler) Update(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "error", "Unauthorized", nil))
	}

	newUser := UserRequest{}
	if errBind := c.Bind(&newUser); errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Error binding data: "+errBind.Error(), nil))
	}

	dataUser := users.User{
		FullName:    newUser.FullName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		PhoneNumber: newUser.PhoneNumber,
		Address:     newUser.Address,
	}

	if errInsert := uh.userService.UpdateProfile(uint(userID), dataUser); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "Failed", "Failed to update account: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "Failed to update account: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "Success", "Successfully updated account", nil))
}

func (uh *UserHandler) Delete(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "error", "Unauthorized", nil))
	}

	if errDelete := uh.userService.DeleteAccount(uint(userID)); errDelete != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "Failed to deleted account: "+errDelete.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "Success", "Succesfully deleted account", nil))
}

func (uh *UserHandler) GetProfile(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "error", "Unauthorized", nil))
	}

	profile, err := uh.userService.GetProfile(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "Failed", "Get user profile failed: "+err.Error(), nil))
	}

	userResponse := UserResponse{
		FullName:    profile.FullName,
		Email:       profile.Email,
		PhoneNumber: profile.PhoneNumber,
		Address:     profile.Address,
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Get user profile successful", echo.Map{"data": userResponse}))
}
