package handlers

import (
	"net/http"
	"restProgram/models"

	"github.com/gin-gonic/gin"
)

type CreatetrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func GetAlltracks(context *gin.Context) {
	var tracks []models.Track
	models.DB.Find(&tracks)

	context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}

func CreateTrack(context *gin.Context) {

	var input CreatetrackInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track := models.Track{Artist: input.Artist, Title: input.Title}
	models.DB.Create(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

func UpdateTrack(context *gin.Context) {
	var track models.Track

	if err := models.DB.Where("id=?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateTrackInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&track).Update(input)
	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

func DeleteTrack(context *gin.Context) {
	var track models.Track

	if err := models.DB.Where("id=?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&track)
	context.JSON(http.StatusOK, gin.H{"tracks": true})
}
