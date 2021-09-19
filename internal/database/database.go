package database

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// ConfigDB db setting
type ConfigDB struct {
	User		string
	Password	string
	Host		string
	Port		string
	Dbname		string
}

var config = ConfigDB{}

var (
	DbClient *gorm.DB
)

// ConnectDB returns init gorm.DB
func init() {
	config.Read()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Jakarta",
		config.Host, config.Port, config.User, config.Dbname, config.Password)
	DbClient, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connection to database")
		log.Fatal(err)
	}
	
	fmt.Println(DbClient)
	fmt.Println("Successfully connected!")
}

// Read and parse the configuration file
func (c *ConfigDB) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}