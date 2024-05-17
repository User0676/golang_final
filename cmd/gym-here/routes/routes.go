package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang_final_project/cmd/gym-here/middlewares"
	"golang_final_project/pkg/gym-here/controllers"
)

type handler struct {
	DB *gorm.DB
}

func Router(r *gin.Engine) {
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	protected.POST("/clients", controllers.CreateClient)
	protected.GET("/clients", controllers.GetClients)
	protected.GET("/clients/:id", controllers.GetClientByID)
	protected.PUT("/clients/:id", controllers.UpdateClient)
	protected.DELETE("/clients/:id", controllers.DeleteClient)
	protected.GET("/clients/instructor/:instructor_id", controllers.GetClientsByInstructor)

	protected.POST("/trainings", controllers.CreateTraining)
	protected.PUT("/trainings/:id", controllers.UpdateTraining)

	protected.POST("/instructors", controllers.CreateInstructor)
	protected.GET("/instructors", controllers.GetInstructors)
	protected.GET("/instructors/:id", controllers.GetInstructorByID)
	protected.PUT("/instructors/:id", controllers.UpdateInstructor)
	protected.DELETE("/instructors/:id", controllers.DeleteInstructor)

	protected.GET("/trainings/:id", controllers.GetWorkoutByID)
}
