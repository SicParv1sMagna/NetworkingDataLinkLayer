package app

import (
	"fmt"
	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/docs"
	ginlogrus "github.com/Toorop/gin-logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run запускает приложение
func (app *Application) Run() {
	const op = "app.Application.Run"
	log := app.log.WithField("operation", op)

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(ginlogrus.Logger(log), gin.Recovery())

	docs.SwaggerInfo.Title = "DataLinkLayer RestAPI"
	docs.SwaggerInfo.Description = "API server for DataLinkLayer application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/api"

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	ApiGroup := router.Group("/api")
	{
		ApiGroup.POST("/channel/code", app.Handler.EncodeSegmentSimulate)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf("%s:%d", app.Config.Host, app.Config.Port)

	log.WithField("addr", addr).Info("HTTP server is running")
	err := router.Run(addr)
	if err != nil {
		log.WithError(err).Error("ошибка запуска HTTP сервера")
	}

	log.Info("Server down")
}
