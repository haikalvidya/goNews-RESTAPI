package service

import (
	db "github.com/haikalvidya/goNews-RESTAPI/internal/go-news/repo"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
)

var ns db.NewsService
var news db.News

func GetNews(id int) (*models.News, error) {
	ns = &news
	// call repo function
	theNews, err := ns.Get(id)
	if err != nil {
		return nil, err
	}

	return theNews, nil
}

func GetAllNews() ([]*models.News, error) {
	ns = &news
	manyNews, err := ns.GetAll()
	if err != nil {
		return nil, err
	}
	return manyNews, nil
}

func AddNews(theNews *db.News) (*db.News, error) {
	ns = theNews
	err := ns.Save()
	if err != nil {
		return err
	}
	return nil
}

func RemoveNews(id int) error {
	ns = &news
	err := ns.Remove(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateNews(param *db.News, id int) error {
	ns = &news
	theNews, err := ns.Get(id)
	if err != nil {
		return err
	}
	err = theNews.Update(param)
	if err != nil {
		return err
	}
	return nil
}

func GetAllNewsByFilter(status string) ([]*models.News, error) {
	ns = &news
	manyNews, err := ns.GetAllByStatus(status)
	if err != nil {
		return nil, err
	}
	return manyNews, nil
}

func GetAllNewsByTopic(theTopic string) ([]*models.News, error) {
	ns = &news
	manyNews, err := ns.GetByTopic(theTopic)
	if err != nil {
		return nil, err
	}
	return manyNews, nil
}