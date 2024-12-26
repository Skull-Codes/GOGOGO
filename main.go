package main

import (
	"fmt"
	"gogogo/db"
	_ "gogogo/docs"
	"gogogo/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.ConnectDB()

	router := gin.Default()

	// Route list:
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.MainRouter(router)
	routes.UserRoutes(router)

	var baseUri string = fmt.Sprintf("%v:%v", os.Getenv("BASE_URI"), os.Getenv("PORT"))
	fmt.Println("Servidor ejecutandose en: " + baseUri)
	server := router.Run(baseUri)

	log.Fatal(server)
}
