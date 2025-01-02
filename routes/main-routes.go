package routes

import (
	"gogogo/controllers"
	"gogogo/middlewares"

	"github.com/gin-gonic/gin"
)

func MainRouter(r *gin.Engine) {
	mainRoutes := r.Group("/")
	{
		mainRoutes.GET("/", func(c *gin.Context) {
			controllers.MainController(c)
		})

		mainRoutes.GET("/media/*uploads", func(c *gin.Context) {
			controllers.UploadServe(c)
		})
	}

	protectedMainRoutes := r.Group("/")
	protectedMainRoutes.Use(middlewares.JwtMiddleware())
	{
		protectedMainRoutes.GET("/test", func(c *gin.Context) {
			controllers.TestAuthController(c)
		})

		/* Descomenta esta ruta y borra la ruta similar para requerir token antes de servir archivos:
		protectedMainRoutes.GET("/:uploads", func(c *gin.Context) {
			controllers.UploadServe(c)
		}) */
	}
}
