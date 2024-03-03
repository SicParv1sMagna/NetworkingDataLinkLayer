package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Run запускает приложение
func (app *Application) Run() {
	r := gin.Default()

	r.POST("/channel/code", app.Handler.EncodeSegmentSimulate)

	addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
	r.Run(addr)
	log.Println("Server down")
}
