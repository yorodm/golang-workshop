package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type data struct {
	Num  int    `json:"Num"`
	Text string `json:"Text"`
}

func create(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "create"})
}

func getOne(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": c.Param("num")})
}

func getAll(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "getAll"})
}

func update(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": c.Param("num")})
}

func delete(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": c.Param("num")})
}

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")

	api.GET("/", getAll)
	api.GET("/:num", getOne)
	api.POST("/", create)
	api.PATCH("/:num", update)
	api.DELETE("/:num", delete)

	router.Run(":8080")
}
