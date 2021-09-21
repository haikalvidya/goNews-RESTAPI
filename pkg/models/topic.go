package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name		string     	`gorm:"not null;" json:"topic"`
	News		[]*News		`gorm:"many2many:news_topics;"`
}

type TopicService interface {
	Get(id int) (*Topic, error)
	GetAll() ([]*Topic, error)
	Save() error
	Remove(id int) error
	Update(*Topic) error
}