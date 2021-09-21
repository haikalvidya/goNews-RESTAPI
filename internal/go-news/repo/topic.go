package repo

import (
	"github.com/haikalvidya/goNews-RESTAPI/internal/database"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
)

// import topic object from models
type Topic models.Topic

// import topic interface from models
type TopicService models.TopicService


func (t Topic) Get(id int) (*models.Topic, error) {
	topic := &models.Topic{}

	err := database.DbClient.Preload("News").First(&topic, id).Error
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (t Topic) GetAll() ([]*models.Topic, error) {
	topics := []*models.Topic{}

	err := database.DbClient.Preload("News").Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (t *Topic) Save() error {
	err := database.DbClient.Create(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Topic) Remove(id int) error {
	err := database.DbClient.First(&t, id).Error
	if err != nil {
		return err
	}
	err = database.DbClient.Delete(&t).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *Topic) Update(topic *models.Topic) error {
	err := database.DbClient.Model(&t).UpdateColumns(Topic{Name: topic.Name}).Error
	if err != nil {
		return err
	}

	return nil
}