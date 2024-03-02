package config

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Server struct {
		Host string
		Port string
	}

	Database struct {
		Host string
		Port string
		Name string
		User string
		Pass string
	}

	JWTSecret string

	Admin struct {
		Email string
	}
}

var app *App

func GetConnection() *gorm.DB {
	conf := GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error(err)
	}

	return db
}

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	config := App{}

	err := godotenv.Load(".env")

	if err != nil {
		log.Error(err)
	}

	config.Server.Host = os.Getenv("APP_HOST")
	config.Server.Port = os.Getenv("APP_PORT")

	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.Name = os.Getenv("DB_NAME")
	config.Database.User = os.Getenv("DB_USER")
	config.Database.Pass = os.Getenv("DB_PASS")

	config.JWTSecret = os.Getenv("JWT_SECRET")

	config.Admin.Email = os.Getenv("ADMIN_EMAIL")

	return &config
}
