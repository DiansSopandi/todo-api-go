// @title Todo API

// @version 1.0

// @description API for managing todo tasks

// @host localhost:8081

// @BasePath /

package main

import (
	"todo-api/docs"
	"todo-api/todo"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.Description = "API for managing todo tasks"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	todo.TodoRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8081")
}
