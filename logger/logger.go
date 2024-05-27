package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func InitLogrus(env string){ 
	
	logrus.SetFormatter(&logrus.TextFormatter{})
	
	switch  env{
	case "debag":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	}
	
	f, err := os.Create("logrus.log")
	if err!=nil{
		logrus.Fatalf("error create log: %s", err.Error())
	}
	multi := io.MultiWriter(f, os.Stdout)
	logrus.SetOutput(multi)
} 