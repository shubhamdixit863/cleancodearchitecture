package main

import (
	"cleanCodego/internal/apis/controllers"
	"cleanCodego/internal/application/services"
	"cleanCodego/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize dependencies
	db := database.NewDatabaseConnection() // Implement this function to return a valid *sql.DB instance
	userRepo := database.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Set up the Gin router
	r := gin.Default()
	registerRoutes(r, userController)
	r.Run(":8080")
}

func registerRoutes(r *gin.Engine, userController *controllers.UserController) {
	users := r.Group("/users")
	{
		users.GET("/:id", userController.GetByID)
		users.POST("", userController.Create)
		users.PUT("", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}
}
