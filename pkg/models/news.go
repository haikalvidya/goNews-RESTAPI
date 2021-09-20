package models

import "gorm.io/gorm"

// news storage model
type News struct {
	gorm.Model
	Title       string     	`gorm:"not null;" json:"title"`
	Author      string     	`gorm:"not null;" json:"author"`
	Content     string     	`gorm:"not null;" json:"content"`
	Status      string		`gorm:"default:draft;not null" json:"status"`
	Slug		string		`gorm:"not null;" json:"slug"`
	Topics      []*Topic    `gorm:"many2many:news_topics;" json:"topics"`
}

type NewsService interface {
	Get(id int) (*News, error)
	GetAll() ([]*News, error)
	GetAllByStatus(status string) ([]*News, error)
	Save(*News) error
	Remove(id int) error
	Update(*News) error
}
