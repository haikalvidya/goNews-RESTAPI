package database

import (
	"log"
	"gorm.io/gorm"
	"github.com/haikalvidya/goNews-RESTAPI/pkg/models"
)

// DBMigrate will create and migrate the tables, then make the some relationships
func DbMigrate() (*gorm.DB, error) {
	conn, err := ConnectDb()
	if err != nil {
		return nil, err
	}
	// defer conn.Close()

	conn.AutoMigrate(&models.News{}, &models.Topic{})
	log.Println("Migration has been processed")

	return conn, nil
}
