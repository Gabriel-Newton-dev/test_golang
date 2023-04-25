package controllers

import (
	"net/http"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/database"
	"github.com/Gabriel-Newton-dev/test_golang/packages/models"
	"github.com/gin-gonic/gin"
)

func DisplayAllCars(c *gin.Context) {
	var Cars []models.Car
	database.DB.Find(&Cars)
	c.JSON(http.StatusOK, Cars)
}

func DisplayCarByID(c *gin.Context) {
	var ID models.Car
	database.DB.First("ID")
	c.JSON(http.StatusOK, ID)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{"API says": "Welcome" + name + "to our API."})
}
