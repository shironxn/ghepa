package main

import (
	"event-planning-app/config"
	"event-planning-app/internal/adapter/handler"
	"event-planning-app/internal/adapter/repository"
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/service"
	"event-planning-app/internal/middleware"
	"event-planning-app/internal/route"
	"event-planning-app/internal/util"

	"github.com/charmbracelet/log"
)

func main() {
	db := config.GetConnection()
	db.AutoMigrate(&domain.User{}, &domain.Event{}, &domain.Comment{})
	log.Info("Database migration successful")

	userRepository := repository.NewUserRepository(db)
	log.Info("User repository created")

	userService := service.NewUserService(userRepository)
	log.Info("User service created")

	userHandler := handler.NewUserHandler(userService)
	log.Info("User handler created")

	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService)

	jwtManager := &util.JWTManager{}
	authMiddleware := middleware.NewAuthMiddleware(*jwtManager)

	route := route.NewRoute(userHandler, eventHandler, *authMiddleware, config.GetConfig())

	route.Initialize()
}
