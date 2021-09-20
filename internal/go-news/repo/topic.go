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

	conn, err := clientDb.ConnectDb()
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	err = conn.Preload("News").First(&topic, id).Error
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (t Topic) GetAll() ([]*models.Topic, error) {
	topics := []*models.Topic{}

	conn, err := clientDb.ConnectDb()
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	err = conn.Preload("News").Find(&topics).Error
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (t *Topic) Save() error {
	conn, err := clientDb.ConnectDb()
	if err != nil {
		return err
	}
	// defer conn.Close()

	err = conn.Create(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Topic) Remove(id int) error {
	conn, err := clientDb.ConnectDb()
	if err != nil {
		return err
	}
	// defer conn.Close()

	err = conn.First(&t, id).Error
	if err != nil {
		return err
	}
	err = conn.Delete(&t).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *Topic) Update(topic *Topic) error {
	conn, err := clientDb.ConnectDb()
	if err != nil {
		return err
	}
	// defer conn.Close()

	err = conn.Model(&t).UpdateColumns(Topic{Name: topic.Name}).Error
	if err != nil {
		return err
	}

	return nil
}