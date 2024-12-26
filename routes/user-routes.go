package routes

import (
	"gogogo/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/api")
	{
		userRoutes.POST("/register", func(c *gin.Context) {
			controllers.RegisterController(c)
		})
		userRoutes.POST("/login", func(c *gin.Context) {
			controllers.LoginController(c)
		})
		userRoutes.POST("/change_password", func(c *gin.Context) {
			controllers.ChangePasswordController(c)
		})
		userRoutes.POST("/delete_user", func(c *gin.Context) {
			controllers.DeleteUserController(c)
		})
	}
}
