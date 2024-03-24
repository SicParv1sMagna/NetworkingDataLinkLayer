package main

import (
	"context"
	"log"

	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/internal/app"
)

// @title DataLinkLayer RestAPI
// @version 1.0
// @description API server for DataLinkLayer application

// @host http://localhost:8081
// @BasePath /api

func main() {
	log.Println("Application start!")
	// создает контекст
	ctx := context.Background()

	// создает Aplication
	application, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// запуск сервера
	application.Run()
	log.Println("Application terminated!")
}
