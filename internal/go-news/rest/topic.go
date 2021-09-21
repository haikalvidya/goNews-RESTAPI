package rest

import (
	"net/http"
	db "github.com/haikalvidya/goNews-RESTAPI/internal/go-news/repo"
	"github.com/haikalvidya/goNews-RESTAPI/internal/go-news/service"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
	"strconv"
	"fmt"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/haikalvidya/goNews-RESTAPI/internal/redis"
)

func CreateTopic(c echo.Context) error {
	topicModel := new(db.Topic)
	// bind request body to the model objects
	if err := c.Bind(topicModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed parsing request body")
	}
	// creating to database with call topic service
	err := service.AddTopic(topicModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to create topic")
	}

	// flush all redis
	err = redis.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}

	return c.JSON(http.StatusCreated, "News created successfully!")
}

func GetTopic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	key := fmt.Sprintf("topic-%d", id)
	if redis.Exists(key) {
		topicCache := redis.Get(key)
		var result models.Topic
		err := json.Unmarshal([]byte(topicCache.(string)), &result)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal topic from redis")
		}
		return c.JSON(http.StatusOK, result)
	}

	response, err := service.GetTopic(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get topic by id")
	}

	// set value of get topic to redis
	err = redis.Set(key, response)
	if err != nil {
		c.Logger().Error("Unable to set in redis cache")
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllTopic(c echo.Context) error {
	key := fmt.Sprintf("topic-all")
	if redis.Exists(key) {
		topicCache := redis.Get(key)
		var result []models.Topic
		err := json.Unmarshal([]byte(topicCache.(string)), &result)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal all topic from redis")
		}
		return c.JSON(http.StatusOK, result)
	}

	response, err := service.GetAllTopic()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get all topic")
	}

	// set value of get topic to redis
	err = redis.Set(key, response)
	if err != nil {
		c.Logger().Error("Unable to set in redis cache")
	}

	return c.JSON(http.StatusOK, response)
}

func RemoveTopic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := service.RemoveTopic(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to remove topic by id")
	}

	key := fmt.Sprintf("topic-%d", id)
	if redis.Exists(key) {
		err := redis.Delete(key)
		if err != nil {
			c.Logger().Error("Failed to delete topic from redis")
		}
	}

	return c.JSON(http.StatusOK, "Topic deleted successfully")
}

func UpdateTopic(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	topicModel := new(models.Topic)
	// bind request body to the model objects
	if err := c.Bind(topicModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed parsing request body")
	}
	// creating to database with call topic service
	err := service.UpdateTopic(topicModel, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update topic")
	}

	// flush all redis
	err = redis.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}

	return c.JSON(http.StatusOK, "Topic updated successfully")
}