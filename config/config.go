package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitEnv(){
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}