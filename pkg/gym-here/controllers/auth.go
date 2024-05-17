package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_final_project/cmd/gym-here/models"
	token "golang_final_project/cmd/gym-here/utils"
	"log"
	"net/http"
)

func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Debug: log the input username
	log.Println("Login attempt for username:", input.Username)

	token, err := models.LoginCheck(input.Username, input.Password)
	if err != nil {
		// Debug: log the error
		log.Println("Login error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password
	u.Role = input.Role

	// Debug: log the username and role
	log.Println("Register attempt:", input.Username, input.Role)

	savedUser, err := u.SaveUser()
	if err != nil {
		// Debug: log the error
		log.Println("Register error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Debug: log the saved user details
	log.Printf("User registered successfully: %+v\n", savedUser)

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
