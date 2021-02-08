package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"problem/1/pkg/handlers"
	"problem/1/pkg/middleware"
	"problem/1/pkg/models"
)

func main() {
	models.InitDB()

	r := gin.Default()
	user := r.Group("")
	user.Use(middleware.TokenAuth())
	// return protobuf format
	user.GET("/users", handlers.GetUserList)
	// return json format
	user.GET("/users/:id", handlers.GetUserDetail)
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
