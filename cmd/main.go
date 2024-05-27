package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"llcmediatelTask/config"
	"llcmediatelTask/logger"
	"llcmediatelTask/internal/handler"
	"llcmediatelTask/internal/server"
	"llcmediatelTask/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	config.InitEnv()
	logger.InitLogrus(os.Getenv("LEVELLOG"))
	logrus.Info(fmt.Sprintf("logging level is %s",os.Getenv("LEVELLOG")))
	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	
	go func() {
		if  err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil{
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Info("http server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	
}