package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("HTTP_PLATFORM_PORT")
		if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
		}
	}

	router.Run(fmt.Sprintf(":%s", port))
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	api.GET("/", getItemsEndpoint)
	api.GET("/:id", getItemEndpoint)
	api.POST("/", postItemEndpoint)
	api.PUT("/", putItemEndpoint)
	api.PATCH("/:id", updateItemEndpoint)
	api.DELETE("/:id", deleteItemEndpoint)

	return router
}

func postItemEndpoint(c *gin.Context) {
	var newItem Item
	c.BindJSON(&newItem)

	repo := createRepository()
	repo.CreateItem(newItem)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

func putItemEndpoint(c *gin.Context) {
	var newItem Item
	c.BindJSON(&newItem)

	repo := createRepository()
	repo.CreateItem(newItem)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

func getItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	repo := createRepository()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, repo.GetItem(id))
}

func getItemsEndpoint(c *gin.Context) {
	repo := createRepository()
	c.JSON(http.StatusOK, repo.GetItems())
}

func updateItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedItem Item
	c.BindJSON(&updatedItem)

	repo := createRepository()
	updatedItem.ID = id
	repo.UpdateItem(updatedItem)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func deleteItemEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	repo := createRepository()
	repo.DeleteItem(id)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func createRepository() TodoRepository {
	repositoryType := os.Getenv("REPOSITORY_TYPE")

	if repositoryType == "Mongo" {
		return &MongoRepository{}
	} else {
		return &MemoryRepository{}
	}
}
