package main

import (
	"go-gin/controllers"
	"go-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
