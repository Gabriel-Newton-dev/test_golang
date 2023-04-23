package routes

import (
	"github.com/Gabriel-Newton-dev/test_golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/healthcheck", controllers.Healthcheck)
	r.Run()
}
