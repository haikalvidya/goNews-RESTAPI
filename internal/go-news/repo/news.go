package repo

import (
	clientDb "github.com/haikalvidya/goNews-RESTAPI/internal/database"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
)

// import news object from models
type News models.News

// import news interface from models
type NewsService models.NewsService


func (c News) Get(id int) (*models.News, error) {
	news := &models.News{}
	if err := clientDb.DbClient.Preload("Topic").First(&news, id).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (c News) GetAll() ([]*models.News, error) {
	news := []*models.News{}
	if err := clientDb.DbClient.Preload("Topic").Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (news *News) Save() error {
	if err := clientDb.DbClient.Save(&news).Error; err != nil {
		return err
	}
	return nil
}

func (c *News) Remove(id int) error {
	tx := clientDb.DbClient.Begin()
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	news := News{}
	if err := tx.First(&news, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	news.Status = "deleted"
	if err := tx.Save(&news).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *News) Update(news *News) error {
	err := clientDb.DbClient.Model(&news).UpdateColumns(News{Title: news.Title, Author: news.Author, Content: news.Content, Status: news.Status, Slug: news.Slug, Topics: news.Topics}).Error
	if err != nil {
		return err
	}
	return nil
}

func (c News) GetAllByStatus(status string) ([]*models.News, error) {
	if status == "deleted" {
		news := []*models.News{}
		err := clientDb.DbClient.Unscoped().Where("status = ?", status).Preload("Topic").Find(&news).Error
		if err != nil {
			return nil, err
		}

		return news, nil
	}

	news := []*models.News{}
	err := clientDb.DbClient.Where("status = ?", status).Preload("Topic").Find(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (c News) GetByTopic(theTopic string) ([]*models.News, error) {
	rows, err := clientDb.DbClient.Raw("SELECT news.id, news.title, news.author, news.content, news.status, news.slug FROM `news_topics` LEFT JOIN news ON news_topics.news_id=news.id WHERE news_topics.topic_id=(SELECT id as topic_id FROM `topics` WHERE topic = ?)", theTopic).Rows() // (*sql.Rows, error)
	defer rows.Close()

	manyNews := make([]*models.News, 0)

	for rows.Next() {
		u := &models.News{}
		err = rows.Scan(&u.ID, &u.Title, &u.Author, &u.Slug, &u.Content, &u.Status)

		if err != nil {
			return nil, err
		}
		manyNews = append(manyNews, u)
	}

	return manyNews, nil
}