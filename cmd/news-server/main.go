package main

import (
	"net/http"
	"github.com/haikalvidya/goNews-RESTAPI/internal/go-news/rest"
	"github.com/haikalvidya/goNews-RESTAPI/internal/redis"
	"github.com/haikalvidya/goNews-RESTAPI/internal/database"
	"log"
	"github.com/labstack/echo/v4"
)

var(
	router = echo.New()
)

func main() {
	mapUrls()

	// init redis 
	redis.InitializeStorage()

	// database migrate
	_, err := database.DbMigrate()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server running at http://localhost:9090/")
	router.Logger.Fatal(router.Start(":9090"))
}

// Define routers
func mapUrls() {

	router.GET("/", index)

	// News
	router.GET("/news", rest.GetAllNews)
	router.GET("/news/:id", rest.GetNews)
	router.GET("/news/?status=:status", rest.GetAllNewsByStatus)
	router.GET("/news/?status=:topic", rest.GetAllNewsByTopic)
	router.POST("/news", rest.CreateNews)
	router.PUT("/news/:id", rest.UpdateNews)
	router.DELETE("/news/:id", rest.RemoveNews)

	// topic
	router.GET("/topic", rest.GetAllTopic)
	router.GET("/topic/:id", rest.GetTopic)
	router.POST("/topic", rest.CreateTopic)
	router.PUT("/topic/:id", rest.UpdateTopic)
	router.DELETE("/topic/:id", rest.RemoveTopic)
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome to news and topic rest api!")
}