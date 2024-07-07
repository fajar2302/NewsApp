package middlewares

import (
	"NEWSAPP/app/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type MiddlewaresInterface interface {
	JWTMiddleware() echo.MiddlewareFunc
	CreateToken(userId int) (string, error)
	ExtractTokenUserId(e echo.Context) int
}

type middlewares struct{}

func NewMiddlewares() MiddlewaresInterface {
	return &middlewares{}
}

func (m *middlewares) JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWT_SECRET),
	})
}

func (m *middlewares) CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     userId,
		"exp":        time.Now().Add(time.Hour * 1).Unix(), // Token expires after 1 hour
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

func (m *middlewares) ExtractTokenUserId(e echo.Context) int {
	header := e.Request().Header.Get("Authorization")
	headerToken := strings.Split(header, " ")
	if len(headerToken) != 2 {
		return 0
	}
	token := headerToken[1]
	tokenJWT, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})

	if err != nil || !tokenJWT.Valid {
		return 0
	}

	claims := tokenJWT.Claims.(jwt.MapClaims)
	userId, isValidUserId := claims["userId"].(float64)
	if !isValidUserId {
		return 0
	}
	return int(userId)
}
