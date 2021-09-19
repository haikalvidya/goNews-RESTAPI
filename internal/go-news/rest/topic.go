package rest

import (
	"fmt"
	"net/http"
	db "github.com/haikalvidya/goNews-RESTAPI/internal/database"
	"github.com/haikalvidya/goNews-RESTAPI/internal/go-news/service"
	"strconv"
	"github.com/labstack/echo/v4"
)

func CreateTopic(c echo.Context) error {
	topicModel := new(db.Topic)
	// bind request body to the model objects
	if err := c.Bind(topicModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed parsing request body")
	}
	// creating to database with call topic service
	response, err := service.AddTopic(topicModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create topic")
	}

	return c.JSON(http.StatusCreated, response)
}

func GetTopic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := service.GetTopic(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get topic by id")
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllTopic(c echo.Context) error {
	response, err := service.GetAllTopic()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get all topic")
	}
	return c.JSON(http.StatusOK, response)
}

func RemoveTopic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := service.RemoveTopic(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to remove topic by id")
	}

	return c.JSON(http.StatusOK, "Topic deleted successfully")
}

func UpdateTopic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	topicModel := new(db.Topic)
	// bind request body to the model objects
	if err := c.Bind(topicModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed parsing request body")
	}
	// creating to database with call topic service
	err := service.UpdateTopic(topicModel, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update topic")
	}

	return c.JSON(http.StatusOK, "Topic updated successfully")
}