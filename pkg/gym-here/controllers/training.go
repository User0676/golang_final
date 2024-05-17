package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_final_project/cmd/gym-here/models"
	"net/http"
	"strconv"
)

type CreateTrainingInput struct {
	Name         string `json:"name" binding:"required"`
	WorkoutTime  string `json:"workout_time" binding:"required"`
	WorkoutDays  string `json:"workout_days" binding:"required"`
	InstructorID uint   `json:"instructor_id" binding:"required"`
}

func CreateTraining(c *gin.Context) {
	var input CreateTrainingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	training := models.Training{
		Name:         input.Name,
		WorkoutTime:  input.WorkoutTime,
		WorkoutDays:  input.WorkoutDays,
		InstructorID: input.InstructorID,
	}

	if err := models.DB.Create(&training).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create training"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "training created successfully", "training": training})
}

type UpdateTrainingInput struct {
	Name        string `json:"name" binding:"required"`
	WorkoutTime string `json:"workout_time" binding:"required"`
	WorkoutDays string `json:"workout_days" binding:"required"`
}

func UpdateTraining(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid training ID"})
		return
	}

	var input UpdateTrainingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var training models.Training
	if err := models.DB.First(&training, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Training not found"})
		return
	}

	training.Name = input.Name
	training.WorkoutTime = input.WorkoutTime
	training.WorkoutDays = input.WorkoutDays

	if err := models.DB.Save(&training).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update training"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Training updated successfully", "training": training})
}
