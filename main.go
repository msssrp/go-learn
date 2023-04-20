package main

import (
	_ "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/msssrp/go-learn/config"
	"github.com/msssrp/go-learn/controllers"
	"github.com/msssrp/go-learn/routes"
)

func getDotEnvVariables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

func main() {

	r := gin.Default()

	DATABASE_URL := getDotEnvVariables("DATABASE_URL")

	client, err := config.Connect(DATABASE_URL)
	if err != nil {
		panic(err)
	}

	userController := controllers.NewUserCollection(client)

	userRouter := r.Group("/api")

	routes.UserRoutes(userRouter, userController)

	r.Run(":8080")
}
