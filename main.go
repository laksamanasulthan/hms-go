package main

import (
	"hms-go/injector"
	"hms-go/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	inj, err := injector.NewInjector()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.RegisterRoutes(router, inj)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	app_port := os.Getenv("APP_PORT")

	addr := ":" + app_port

	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
