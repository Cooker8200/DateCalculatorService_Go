package routes

import (
	aws "DateCalculator-Service-Go/data/aws"
	mongo "DateCalculator-Service-Go/data/mongo"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.GET("/dates/aws", aws.GetAllDates)
	router.PUT("/dates/aws", aws.AddNewDate)
	router.DELETE("/dates/aws", aws.RemoveDate)

	router.GET("/dates/mongo", mongo.GetAllDates)
	router.PUT("/dates/mongo")
	router.DELETE("/dates/mongo")

	router.POST("/dates/monogo/populate")
	router.DELETE("/dates/mongo/wipe")

	router.Run("localhost:3001")
}
