package models

import "github.com/jinzhu/gorm"

type Topic struct {
	gorm.Model
	Name		string     	`gorm:"not null;" json:"topic"`
	News		[]News		`gorm:"many2many:news_topics;`
}

type TopicService interface {
	Get(id int) (*Topic, error)
	GetAll() ([]Topic, error)
	Save(*Topic) error
	Remove(id int) error
	Update(*Topic) error
}