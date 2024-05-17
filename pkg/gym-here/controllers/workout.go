package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang_final_project/cmd/gym-here/models"
)

type ClientInfo struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	TrainingID uint   `json:"training_id"`
}

type InstructorInfo struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	ProfileSport   string  `json:"profile_sport"`
	Qualification  *string `json:"qualification"`
	WorkExperience *int    `json:"work_experience"`
}

type WorkoutResponse struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	WorkoutTime string         `json:"workout_time"`
	WorkoutDays string         `json:"workout_days"`
	Instructor  InstructorInfo `json:"instructor"`
	Clients     []ClientInfo   `json:"clients"`
}

func GetWorkoutByID(c *gin.Context) {
	workoutID := c.Param("id")

	var workout models.Training
	if err := models.DB.Preload("Instructor").Preload("Clients").First(&workout, workoutID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout not found"})
		return
	}

	var clientInfos []ClientInfo
	for _, client := range workout.Clients {
		clientInfos = append(clientInfos, ClientInfo{
			ID:         client.ID,
			Name:       client.Name,
			Age:        client.Age,
			TrainingID: client.TrainingID,
		})
	}

	instructorInfo := InstructorInfo{
		ID:             workout.Instructor.ID,
		Name:           workout.Instructor.Name,
		ProfileSport:   workout.Instructor.ProfileSport,
		Qualification:  workout.Instructor.Qualification,
		WorkExperience: workout.Instructor.WorkExperience,
	}

	workoutResponse := WorkoutResponse{
		ID:          workout.ID,
		Name:        workout.Name,
		WorkoutTime: workout.WorkoutTime,
		WorkoutDays: workout.WorkoutDays,
		Instructor:  instructorInfo,
		Clients:     clientInfos,
	}

	c.JSON(http.StatusOK, gin.H{"workout": workoutResponse})
}
