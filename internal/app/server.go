package app

import (
	"fmt"
	"log"

	"github.com/SicParv1sMagna/NetworkingDataLinkLayer/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run запускает приложение
func (app *Application) Run() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "DataLinkLayer RestAPI"
	docs.SwaggerInfo.Description = "API server for DataLinkLayer application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/api"

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	
	ApiGroup := r.Group("/api") 
	{
		ApiGroup.POST("/channel/code", app.Handler.EncodeSegmentSimulate)
	} 
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
	r.Run(addr)
	
	log.Println("Server down")
}
