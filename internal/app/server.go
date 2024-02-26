package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Run запускает приложение
func (app *Application) Run() {
	r := gin.Default()

	// эндпоинт получения всех монет
	//r.POST("/api/v1/segment", app.Handler.GetAllCoins)

	addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
	r.Run(addr)
	log.Println("Server down")
}
