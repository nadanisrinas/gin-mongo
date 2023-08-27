package main

import (
	"sesi7-challenge/controllers"

	_ "sesi7-challenge/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title Post API
//@version 1.0.0
//@description service for managing posts
//@termsOfService http://swagger.io/terms/
//@contact.name API support
//@contact.email nada@gmail.com
//@lisence.name Apache 2.0
//@host localhost:3000
//@BasePath /

func main() {
	router := gin.Default()
	//Read all
	router.POST("/", controllers.CreatePost)
	//Read one
	// called as localhost:3000/getOne/{id}
	router.GET("getOne/:postId", controllers.ReadOnePost)
	//Update one
	// called as localhost:3000/update/{id}
	router.PUT("/update/:postId", controllers.UpdatePost)
	//Delete one
	// called as localhost:3000/delete/{id}
	router.DELETE("/delete/:postId", controllers.DeletePost)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost: 3000")
}
