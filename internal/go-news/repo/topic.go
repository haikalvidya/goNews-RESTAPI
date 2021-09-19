package repo

import (
	clientDb "github.com/haikalvidya/goNews-RESTAPI/internal/database"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
)

// import topic object from models
type Topic models.Topic

// import topic interface from models
type TopicService models.TopicService


func (c Topic) Get(id int) (*Topic, error) {
	topic := &Topic{}
	err := clientDb.DbClient.Preload("News").First(&topic, id).Error
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (c Topic) GetAll() ([]Topic, error) {
	topics := []Topic{}
	err := clientDb.DbClient.Preload("News").Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (c *Topic) Save(topic *Topic) error {
	err := clientDb.DbClient.Save(&topic).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Topic) Remove(id int) error {
	topic := &Topic{}
	err := clientDb.DbClient.First(&topic, id).Error
	if err != nil {
		return err
	}
	err = clientDb.DbClient.Delete(&topic).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Topic) Update(topic *Topic) error {
	err := clientDb.DbClient.Model(&topic).UpdateColumns(Topic{Name: topic.Name}).Error
	if err != nil {
		return err
	}

	return nil
}