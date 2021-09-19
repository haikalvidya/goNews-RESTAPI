package rest

import (
	"net/http"
	db "github.com/haikalvidya/goNews-RESTAPI/internal/go-news/repo"
	"github.com/haikalvidya/goNews-RESTAPI/internal/go-news/service"
	"strconv"
	"github.com/labstack/echo/v4"
)

func CreateNews(c echo.Context) error {
	newsModel := new(db.News)
	// bind request body to the model objects
	if err := c.Bind(newsModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed parsing request body")
	}
	// creating to database with call news service
	err := service.AddNews(newsModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create news")
	}

	return c.JSON(http.StatusCreated, "News created successfully!")
}

func GetNews(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := service.GetNews(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get news by id")
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllNews(c echo.Context) error {
	response, err := service.GetAllNews()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get all news")
	}
	return c.JSON(http.StatusOK, response)
}

func GetAllNewsByStatus(c echo.Context) error {
	status := c.Param("status")
	
	response, err := service.GetAllNewsByFilter(status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get news by status") 
	}
	return c.JSON(http.StatusOK, response)
}

func GetAllNewsByTopic(c echo.Context) error{
	topic := c.Param("topic")
	
	response, err := service.GetAllNewsByTopic(topic)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get news by status") 
	}
	return c.JSON(http.StatusOK, response)
}

func RemoveNews(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.RemoveNews(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to remove news by id")
	}

	return c.JSON(http.StatusOK, "News deleted successfully")
}

func UpdateNews(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	newsModel := new(db.News)
	// bind request body to the model objects
	if err := c.Bind(newsModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed parsing request body")
	}
	// creating to database with call news service
	err := service.UpdateNews(newsModel, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update news")
	}

	return c.JSON(http.StatusOK, "News updated successfully")
}