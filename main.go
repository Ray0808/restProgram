package main

import (
	"restProgram/handlers"
	"restProgram/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()

	route.GET("/tracks", handlers.GetAlltracks)
	route.POST("/createTrack", handlers.CreateTrack)
	route.PUT("/update/:id", handlers.UpdateTrack)
	route.DELETE("/delete/:id", handlers.DeleteTrack)

	route.Run()
}
