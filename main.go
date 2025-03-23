package main

import (
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	public := r.Group("/api")
	// public.POST("/register", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "this is register entrypoint"})
	// })

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}
