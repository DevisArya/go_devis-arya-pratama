package config

import (
	"code-competence/models"
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	DB *gorm.DB
)

func Init() {
	initDB()
	InitMigration()
}

type Config struct {
	Username string
	Password string
	Port     string
	Host     string
	Name     string
}

func initDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

func InitMigration() {
	DB.AutoMigrate(
		&models.User{},
		&models.Items{},
		&models.Category{},
	)
}
