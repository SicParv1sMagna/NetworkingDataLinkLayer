package main

import (
	"context"
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/lib/logger/handlers/logruspretty"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/app"
)

// @title DataLinkLayer RestAPI
// @version 1.0
// @description API server for DataLinkLayer application

// @host http://localhost:8081
// @BasePath /api

func main() {
	// инициализирует логгер
	log := setupLogger()

	log.Info("Application start")
	// создает контекст
	ctx := context.Background()

	// создает Aplication
	application, err := app.New(ctx, log)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// запуск сервера
	application.Run()
	log.Info("Application terminated")
}

func setupLogger() *logrus.Logger {
	var log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	return setupPrettyLogrus(log)
}

func setupPrettyLogrus(log *logrus.Logger) *logrus.Logger {
	prettyHandler := logruspretty.NewPrettyHandler(os.Stdout)
	log.SetFormatter(prettyHandler)
	return log
}
