package config

import (
	"os"
	"strconv"
)

var (
	JWT_SECRET string
)

type AppConfig struct {
	USER     string
	PASSWORD string
	HOST     string
	PORT     int
	DBNAME   string
}

func ReadEnv() *AppConfig {
	var app = AppConfig{}
	app.USER = os.Getenv("DBUSER")
	app.PASSWORD = os.Getenv("DBPASS")
	app.HOST = os.Getenv("DBHOST")
	portConv, errConv := strconv.Atoi(os.Getenv("DBPORT"))
	if errConv != nil {
		panic("error convert dbport")
	}
	app.PORT = portConv
	app.DBNAME = os.Getenv("DBNAME")
	JWT_SECRET = os.Getenv("JWTSECRET")
	return &app
}

func InitConfig() *AppConfig {
	return ReadEnv()
}
