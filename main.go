package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kritimauludin/go-restapi-gin-gorm-mysql/controllers/productcontroller"
	"github.com/kritimauludin/go-restapi-gin-gorm-mysql/models"
)

func main()  {
	routes := gin.Default()
	models.ConnectDatabase()

	routes.GET("/api/products", productcontroller.Index)
	routes.GET("/api/product/:id", productcontroller.Show)
	routes.POST("/api/product", productcontroller.Create)
	routes.PUT("/api/product/:id", productcontroller.Update)
	routes.DELETE("/api/product", productcontroller.Delete)


	routes.Run()
}