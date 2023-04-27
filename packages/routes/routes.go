package routes

import (
	"github.com/Gabriel-Newton-dev/test_golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/healthcheck", controllers.Healthcheck)
	r.GET("/cars", controllers.DisplayAllCars)
	r.GET("/cars/:id", controllers.DisplayCarByID)
	r.GET("/:name", controllers.Salutation)
	r.POST("/cars", controllers.CreateCardRegister)
	r.Run()
}
