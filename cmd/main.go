package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ArikuWoW/extract/initializers"
	"github.com/ArikuWoW/extract/models"
	"github.com/ArikuWoW/extract/pkg/handler"
	"github.com/ArikuWoW/extract/pkg/repository"
	"github.com/ArikuWoW/extract/pkg/service"
	"github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	db := initializers.ConnectDB()

	repos := repository.NewRepository(db)

	service := service.NewService(repos)

	handlers := handler.NewHandler(service)

	srv := new(models.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on db conection close: %s", err.Error())
	}
}
