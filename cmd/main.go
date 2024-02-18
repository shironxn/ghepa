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
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	jwtManager := &util.JWTManager{}
	authMiddleware := middleware.NewAuthMiddleware(*jwtManager)

	route := route.NewRoute(userHandler, eventHandler, commentHandler, *authMiddleware, config.GetConfig())

	route.Initialize()
}
