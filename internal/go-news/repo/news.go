package repo

import (
	clientDb "github.com/haikalvidya/goNews-RESTAPI/internal/database"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
)

// import news object from models
type News models.News

// import news interface from models
// type NewsService models.NewsService


func (n News) Get(id int) (*models.News, error) {
	// news := &models.News{}

	conn, err := clientDb.ConnectDb()
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	if err = conn.Preload("Topics").First(&n, id).Error; err != nil {
		return nil, err
	}
	news := models.News(n)
	return &news, nil
}

func (n News) GetAll() ([]*models.News, error) {
	manyNews := []*models.News{}

	conn, err := clientDb.ConnectDb()
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	if err = conn.Preload("Topics").Find(&manyNews).Error; err != nil {
		return nil, err
	}
	return manyNews, nil
}

func (news *News) Save() error {

	conn, err := clientDb.ConnectDb()
	if err != nil {
		return err
	}
	// defer conn.Close()

	if err = conn.Create(&news).Error; err != nil {
		return err
	}

	return nil
}

func (n *News) Remove(id int) error {

	conn, err := clientDb.ConnectDb()
	if err != nil {
		return err
	}
	// defer conn.Close()

	tx := conn.Begin()
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	news := News{}
	if err := tx.First(&n, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	news.Status = "deleted"
	if err := tx.Save(&n).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (n *News) Update(news *News) error {

	conn, err := clientDb.ConnectDb()
	if err != nil {
		return err
	}
	// defer conn.Close()

	err = conn.Model(&n).UpdateColumns(News{Title: news.Title, Author: news.Author, Content: news.Content, Status: news.Status, Slug: news.Slug, Topics: news.Topics}).Error
	if err != nil {
		return err
	}
	return nil
}

func (n News) GetAllByStatus(status string) ([]*models.News, error) {

	conn, err := clientDb.ConnectDb()
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	if status == "deleted" {
		news := []*models.News{}
		err := conn.Unscoped().Where("status = ?", status).Preload("Topics").Find(&news).Error
		if err != nil {
			return nil, err
		}

		return news, nil
	}

	news := []*models.News{}
	err = conn.Where("status = ?", status).Preload("Topics").Find(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (n News) GetByTopic(theTopic string) ([]*models.News, error) {
	conn, err := clientDb.ConnectDb()
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	rows, err := conn.Raw("SELECT news.id, news.title, news.author, news.content, news.status, news.slug FROM `news_topics` LEFT JOIN news ON news_topics.news_id=news.id WHERE news_topics.topic_id=(SELECT id as topic_id FROM `topics` WHERE topic = ?)", theTopic).Rows() // (*sql.Rows, error)
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