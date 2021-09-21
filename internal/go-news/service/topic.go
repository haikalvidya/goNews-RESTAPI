package service

import (
	db "github.com/haikalvidya/goNews-RESTAPI/internal/go-news/repo"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
)

var topicService db.TopicService
var topic db.Topic

func GetTopic(id int) (*models.Topic, error) {
	topicService = &topic
	// call repo function
	theTopic, err := topicService.Get(id)
	if err != nil {
		return nil, err
	}

	return theTopic, nil
}

func GetAllTopic() ([]*models.Topic, error) {
	topicService = &topic
	manyTopic, err := topicService.GetAll()
	if err != nil {
		return nil, err
	}
	return manyTopic, nil
}

func AddTopic(theTopic *db.Topic) error {
	topicService = theTopic
	err := topicService.Save()
	if err != nil {
		return err
	}
	return nil
}

func RemoveTopic(id int) error {
	topicService = &topic
	err := topicService.Remove(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTopic(param *models.Topic, id int) error {
	topicService := &topic
	param.ID = uint(id)
	err := topicService.Update(param)
	if err != nil {
		return err
	}
	return nil
}