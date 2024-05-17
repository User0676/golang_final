package main

import (
	"github.com/gin-gonic/gin"
	"golang_final_project/cmd/gym-here/models"
	"golang_final_project/cmd/gym-here/routes"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()
	routes.Router(r)
	r.Run(":8080")

}
