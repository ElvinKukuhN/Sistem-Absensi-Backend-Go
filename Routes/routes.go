package Routes

import (
	"Sistem-Absensi-Backend-Go/Controller/AuthController"
	"Sistem-Absensi-Backend-Go/Controller/PermissionController"
	"Sistem-Absensi-Backend-Go/Controller/RoleController"
	"Sistem-Absensi-Backend-Go/Controller/UserController"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	// Auth routes
	auth := api.Group("/auth")
	{
		auth.POST("/signup", AuthController.SignUpHandler)
		auth.POST("/signin", AuthController.SignInHandler)
	}

	// User routes
	user := api.Group("/user")
	{
		user.POST("/create", UserController.CreateHandler)
		user.GET("", UserController.GetHandler)
		user.PUT("/:id", UserController.UpdateHandler)
		//user.DELETE("", UserController.DeleteHandler)
	}

	// Role routes
	role := api.Group("/role")
	{
		role.POST("/create", RoleController.CreateHandler)
		role.GET("", RoleController.GetHandler)
		role.DELETE("", RoleController.DeleteHandler)
	}

	permission := api.Group("/permission")
	{
		permission.POST("/create", PermissionController.CreateHandler)
		permission.GET("", PermissionController.GetHandler)
		permission.DELETE("", PermissionController.DeleteHandler)
	}

}
