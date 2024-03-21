package main

import (
	"fmt"

	"github.com/9500073161/skill-map-prod/handlers"
	"github.com/9500073161/skill-map-prod/managers"
	"github.com/9500073161/skill-map-prod/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	storage.InitializeDatabase()
}

func main() {
	fmt.Println("Skill Map-Prod-v1")

	router := gin.Default()

	router.Use(cors.Default())

	router.Use(static.Serve("/", static.LocalFile("frontend/build", false)))
	router.Use(static.Serve("/users", static.LocalFile("frontend/build", false)))
	router.Use(static.Serve("/skills", static.LocalFile("frontend/build", false)))

	userManger := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManger)
	userHandler.RegisterUserApis(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
