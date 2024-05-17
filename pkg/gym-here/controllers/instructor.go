package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang_final_project/cmd/gym-here/models"
)

type InstructorResponse struct {
	ID             uint    `json:"id"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	Name           string  `json:"name"`
	ProfileSport   string  `json:"profile_sport"`
	Qualification  *string `json:"qualification"`
	WorkExperience *int    `json:"work_experience"`
}

// CreateInstructorInput defines the input structure for creating an instructor
type CreateInstructorInput struct {
	Name           string  `json:"name" binding:"required"`
	ProfileSport   string  `json:"profile_sport" binding:"required"`
	Qualification  *string `json:"qualification"`
	WorkExperience *int    `json:"work_experience"`
}

// UpdateInstructorInput defines the input structure for updating an instructor
type UpdateInstructorInput struct {
	Name           string  `json:"name" binding:"required"`
	ProfileSport   string  `json:"profile_sport" binding:"required"`
	Qualification  *string `json:"qualification"`
	WorkExperience *int    `json:"work_experience"`
}

// CreateInstructor creates a new instructor
func CreateInstructor(c *gin.Context) {
	var input CreateInstructorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructor := models.Instructor{
		Name:           input.Name,
		ProfileSport:   input.ProfileSport,
		Qualification:  input.Qualification,
		WorkExperience: input.WorkExperience,
	}

	if err := models.DB.Create(&instructor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create instructor", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Instructor created successfully", "instructor": instructor})
}

// GetInstructors retrieves all instructors
func GetInstructors(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	sortBy := c.DefaultQuery("sortBy", "created_at")
	order := c.DefaultQuery("order", "asc")

	profileSport := c.Query("profile_sport")

	var instructors []models.Instructor
	query := models.DB.Offset(offset).Limit(pageSize).Order(sortBy + " " + order)

	if profileSport != "" {
		query = query.Where("profile_sport = ?", profileSport)
	}

	if err := query.Find(&instructors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var instructorResponses []InstructorResponse
	for _, instructor := range instructors {
		instructorResponses = append(instructorResponses, InstructorResponse{
			ID:             instructor.ID,
			CreatedAt:      instructor.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:      instructor.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Name:           instructor.Name,
			ProfileSport:   instructor.ProfileSport,
			Qualification:  instructor.Qualification,
			WorkExperience: instructor.WorkExperience,
		})
	}

	c.JSON(http.StatusOK, gin.H{"instructors": instructorResponses})
}

// GetInstructorByID retrieves a specific instructor by ID
func GetInstructorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructor ID"})
		return
	}

	var instructor models.Instructor
	if err := models.DB.First(&instructor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"instructor": instructor})
}

// UpdateInstructor updates an existing instructor
func UpdateInstructor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructor ID"})
		return
	}

	var input UpdateInstructorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var instructor models.Instructor
	if err := models.DB.First(&instructor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}

	instructor.Name = input.Name
	instructor.ProfileSport = input.ProfileSport
	instructor.Qualification = input.Qualification
	instructor.WorkExperience = input.WorkExperience

	if err := models.DB.Save(&instructor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update instructor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Instructor updated successfully", "instructor": instructor})
}

// DeleteInstructor deletes an instructor by ID
func DeleteInstructor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instructor ID"})
		return
	}

	var instructor models.Instructor
	if err := models.DB.First(&instructor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		return
	}

	if err := models.DB.Delete(&instructor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete instructor"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Instructor deleted successfully"})
}
