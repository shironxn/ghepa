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
		Name     string
		Email    string
		Password string
	}
}

var app *App

func GetConnection() *gorm.DB {
	conf := GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Database.User, conf.Database.Pass, conf.Database.Host, conf.Database.Port, conf.Database.Name)
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
	conf := App{}

	err := godotenv.Load(".env")

	if err != nil {
		log.Error(err)
	}

	conf.Server.Host = os.Getenv("APP_HOST")
	conf.Server.Port = os.Getenv("APP_PORT")

	conf.Database.Host = os.Getenv("DB_HOST")
	conf.Database.Port = os.Getenv("DB_PORT")
	conf.Database.Name = os.Getenv("DB_NAME")
	conf.Database.User = os.Getenv("DB_USER")
	conf.Database.Pass = os.Getenv("DB_PASS")

	conf.JWTSecret = os.Getenv("JWT_SECRET")

	conf.Admin.Name = os.Getenv("ADMIN_NAME")
	conf.Admin.Email = os.Getenv("ADMIN_EMAIL")
	conf.Admin.Password = os.Getenv("ADMIN_PASSWORD")

	return &conf
}
