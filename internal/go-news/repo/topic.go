package repo

import (
	clientDb "github.com/haikalvidya/goNews-RESTAPI/internal/database"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
)

// import topic object from models
type Topic models.Topic

// import topic interface from models
// type TopicService models.TopicService


func (t Topic) Get(id int) (*models.Topic, error) {
	topic := &models.Topic{}
	err := clientDb.DbClient.Preload("News").First(&topic, id).Error
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (t Topic) GetAll() ([]*models.Topic, error) {
	topics := []*models.Topic{}
	err := clientDb.DbClient.Preload("News").Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (t *Topic) Save() error {
	err := clientDb.DbClient.Save(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Topic) Remove(id int) error {
	// topic := &Topic{}
	err := clientDb.DbClient.First(&t, id).Error
	if err != nil {
		return err
	}
	err = clientDb.DbClient.Delete(&t).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *Topic) Update(topic *Topic) error {
	err := clientDb.DbClient.Model(&t).UpdateColumns(Topic{Name: topic.Name}).Error
	if err != nil {
		return err
	}

	return nil
}