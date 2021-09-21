package database

import (
	"fmt"
	"log"

	// "github.com/BurntSushi/toml"
	"gorm.io/gorm"
	// "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

// ConfigDB db setting
// type ConfigDB struct {
// 	User		string
// 	Password	string
// 	Host		string
// 	Port		string
// 	Dbname		string
// }

// var config = ConfigDB{}

var (
	DbClient *gorm.DB
	err error
)

// ConnectDb returns init gorm.DB
// func ConnectDb() (*gorm.DB, error) {
func InitDb() {
	// config.Read()

	// connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Jakarta",
	// 	config.Host, config.Port, config.User, config.Dbname, config.Password)
	// DbClient, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	DbClient, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connection to database")
		log.Fatal(err)
	}
	
	fmt.Println(DbClient)
	fmt.Println("Successfully connected!")

	// return DbClient, nil
}

// Read and parse the configuration file
// func (c *ConfigDB) Read() {
// 	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
// 		log.Fatal(err)
// 	}
// }