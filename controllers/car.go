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
	var car models.Car
	id := c.Params.ByName("id")
	database.DB.First(&car, id)
	if car.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Car not found",
		})
	}
	c.JSON(http.StatusOK, car)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{"API says": "Welcome " + name + " to our API."})
}

func CreateCardRegister(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}
	database.DB.Create(&models.Cars)
	c.JSON(http.StatusOK, car)

}
