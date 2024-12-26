package controllers

import (
	"github.com/gin-gonic/gin"
)

// @Summary Muestra que la api está funcionando
// @Description Endpoint para mostrar un mensaje de petición correcta
// @Tags Main
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string{message=string} "Peticion correcta" example({"message": "Corrected request get to Main"})
// @Router / [get]
func MainController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Corrected request get to Main",
	})
}

// @Summary Test auth
// @Description Este endpoint requiere un Bearer Token para autenticarse.
// @Security BearerAuth
// @Tags Main
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} map[string]string{message=string} "Ruta test auth correcta" example({"message": "Test auth uri correct"})
// @Router /test [get]
func TestAuthController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Test auth uri correct",
	})
}
