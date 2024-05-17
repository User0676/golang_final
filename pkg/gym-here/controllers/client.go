package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_final_project/cmd/gym-here/models"
	"net/http"
	"strconv"
)

type ClientResponse struct {
	ID          uint   `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	TrainingID  uint   `json:"training_id"`
	WorkoutTime string `json:"workout_time"`
	WorkoutDays string `json:"workout_days"`
}
type CustomClientResponse struct {
	ID         uint    `json:"id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at,omitempty"`
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	TrainingID uint    `json:"training_id"`
}
type CreateClientInput struct {
	Name       string `json:"name" binding:"required"`
	Age        int    `json:"age" binding:"required"`
	TrainingID uint   `json:"training_id" binding:"required"`
}
type UpdateClientInput struct {
	Name       string `json:"name" binding:"required"`
	Age        int    `json:"age" binding:"required"`
	TrainingID uint   `json:"training_id" binding:"required"`
}

type ClientInput struct {
	Name       string `json:"name" binding:"required"`
	Age        int    `json:"age" binding:"required,min=0"`
	TrainingID uint   `json:"training_id" binding:"required"`
}

func CreateClient(c *gin.Context) {
	var input CreateClientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := models.Client{
		Name:       input.Name,
		Age:        input.Age,
		TrainingID: input.TrainingID,
	}

	if err := models.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client"})
		return
	}

	// Load the associated Training entity
	models.DB.Preload("Training.Instructor").First(&client, client.ID)

	c.JSON(http.StatusOK, gin.H{"message": "client created successfully", "client": client})
}

func GetClientsByInstructor(c *gin.Context) {
	instructorID, err := strconv.Atoi(c.Param("instructor_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructor ID"})
		return
	}

	var clients []models.Client
	if err := models.DB.Where("training_id IN (SELECT id FROM trainings WHERE instructor_id = ?)", instructorID).Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var clientResponses []ClientResponse
	for _, client := range clients {
		var training models.Training
		if err := models.DB.First(&training, client.TrainingID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		clientResponses = append(clientResponses, ClientResponse{
			ID:          client.ID,
			CreatedAt:   client.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   client.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Name:        client.Name,
			Age:         client.Age,
			TrainingID:  client.TrainingID,
			WorkoutTime: training.WorkoutTime,
			WorkoutDays: training.WorkoutDays,
		})
	}

	c.JSON(http.StatusOK, gin.H{"clients": clientResponses})
}

func GetClients(c *gin.Context) {
	// Default values for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	// Sorting parameters
	sortBy := c.DefaultQuery("sortBy", "created_at")
	order := c.DefaultQuery("order", "asc")

	// Filtering parameters
	age := c.Query("age")
	trainingID := c.Query("training_id")

	var clients []models.Client
	query := models.DB.Offset(offset).Limit(pageSize).Order(sortBy + " " + order)

	if age != "" {
		query = query.Where("age = ?", age)
	}
	if trainingID != "" {
		query = query.Where("training_id = ?", trainingID)
	}

	if err := query.Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var clientResponses []CustomClientResponse
	for _, client := range clients {
		clientResponse := CustomClientResponse{
			ID:         client.ID,
			CreatedAt:  client.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:  client.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Name:       client.Name,
			Age:        client.Age,
			TrainingID: client.TrainingID,
		}
		if client.DeletedAt != nil {
			deletedAt := client.DeletedAt.Format("2006-01-02T15:04:05Z07:00")
			clientResponse.DeletedAt = &deletedAt
		}
		clientResponses = append(clientResponses, clientResponse)
	}

	c.JSON(http.StatusOK, gin.H{"clients": clientResponses})
}

func GetClientByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	client, err := models.GetClientByID(models.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"client": client})
}

func UpdateClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	var input UpdateClientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client models.Client
	if err := models.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	client.Name = input.Name
	client.Age = input.Age
	client.TrainingID = input.TrainingID

	if err := models.DB.Save(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client"})
		return
	}

	// Load the associated Training entity
	models.DB.Preload("Training").First(&client, client.ID)

	c.JSON(http.StatusOK, gin.H{"message": "Client updated successfully", "client": client})
}

func DeleteClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	if err := models.DeleteClient(models.DB, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "client deleted successfully"})
}
