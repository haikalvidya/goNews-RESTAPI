package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	db "github.com/haikalvidya/goNews-RESTAPI/internal/go-news/repo"
	"github.com/haikalvidya/goNews-RESTAPI/internal/go-news/service"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/haikalvidya/goNews-RESTAPI/internal/redis"
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
	
	// flush all redis
	err = redis.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}

	return c.JSON(http.StatusCreated, "News created successfully!")
}

func GetNews(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	key := fmt.Sprintf("news-%d", id)
	if redis.Exists(key) {
		newsCache := redis.Get(key)
		var result models.News
		err := json.Unmarshal([]byte(newsCache.(string)), &result)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal news from redis")
		}
		return c.JSON(http.StatusOK, result)
	}

	response, err := service.GetNews(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get news by id")
	}

	// set value of get news to redis
	err = redis.Set(key, response)
	if err != nil {
		c.Logger().Error("Unable to set in redis cache")
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllNews(c echo.Context) error {
	key := fmt.Sprintf("news-all")
	if redis.Exists(key) {
		newsCache := redis.Get(key)
		var result []models.News
		err := json.Unmarshal([]byte(newsCache.(string)), &result)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal all news from redis")
		}
		return c.JSON(http.StatusOK, result)
	}
	response, err := service.GetAllNews()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get all news")
	}

	// set value of get news to redis
	err = redis.Set(key, response)
	if err != nil {
		c.Logger().Error("Unable to set in redis cache")
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllNewsByStatus(c echo.Context) error {
	status := c.Param("status")
	
	key := fmt.Sprintf("news-%s", status)
	if redis.Exists(key) {
		newsCache := redis.Get(key)
		var result []models.News
		err := json.Unmarshal([]byte(newsCache.(string)), &result)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal all news by status from redis")
		}
		return c.JSON(http.StatusOK, result)
	}

	response, err := service.GetAllNewsByFilter(status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get news by status") 
	}

	// set value of get news to redis
	err = redis.Set(key, response)
	if err != nil {
		c.Logger().Error("Unable to set in redis cache")
	}

	return c.JSON(http.StatusOK, response)
}

func GetAllNewsByTopic(c echo.Context) error{
	topic := c.Param("topic")

	key := fmt.Sprintf("news-%s", topic)
	if redis.Exists(key) {
		newsCache := redis.Get(key)
		var result []models.News
		err := json.Unmarshal([]byte(newsCache.(string)), &result)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to unmarshal all news by topic from redis")
		}
		return c.JSON(http.StatusOK, result)
	}
	
	response, err := service.GetAllNewsByTopic(topic)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get news by status") 
	}

	// set value of get news to redis
	err = redis.Set(key, response)
	if err != nil {
		c.Logger().Error("Unable to set in redis cache")
	}

	return c.JSON(http.StatusOK, response)
}

func RemoveNews(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.RemoveNews(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to remove news by id")
	}

	key := fmt.Sprintf("news-%d", id)
	if redis.Exists(key) {
		err := redis.Delete(key)
		if err != nil {
			c.Logger().Error("Failed to delete news from redis")
		}
	}

	return c.JSON(http.StatusOK, "News deleted successfully")
}

func UpdateNews(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	newsModel := new(models.News)
	// bind request body to the model objects
	if err := c.Bind(newsModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed parsing request body")
	}
	
	// creating to database with call news service
	err := service.UpdateNews(newsModel, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to update news")
	}

	// flush all redis
	err = redis.Flush()
	if err != nil {
		c.Logger().Error("unable to flush redis cache")
	}

	return c.JSON(http.StatusOK, "News updated successfully")
}