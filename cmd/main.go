package main

import (
	"ghepa/config"
	"ghepa/internal/adapter/handler"
	"ghepa/internal/adapter/repository"
	"ghepa/internal/core/domain"
	"ghepa/internal/core/service"
	"ghepa/internal/middleware"
	"ghepa/internal/route"
	"ghepa/internal/util"
)

func main() {
	db := config.GetConnection()
	db.AutoMigrate(&domain.User{}, &domain.Event{}, &domain.Comment{})

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
