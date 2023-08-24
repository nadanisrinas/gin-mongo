package main

import (
	"sesi7-challenge/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/", controllers.CreatePost)

	// called as localhost:3000/getOne/{id}
	router.GET("getOne/:postId", controllers.ReadOnePost)

	// called as localhost:3000/update/{id}
	router.PUT("/update/:postId", controllers.UpdatePost)

	// called as localhost:3000/delete/{id}
	router.DELETE("/delete/:postId", controllers.DeletePost)

	router.Run("localhost: 3000")
}
