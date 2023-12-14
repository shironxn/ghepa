package route

import (
	"event-planning-app/config"
	"event-planning-app/internal/core/port"
	"event-planning-app/internal/middleware"
	"event-planning-app/internal/util"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
)

type Route struct {
	user           port.UserHandler
	response       util.Response
	authMiddleware middleware.AuthMiddleware
	config         *config.App
}

func NewRoute(user port.UserHandler, auth middleware.AuthMiddleware, config *config.App) *Route {
	return &Route{
		user:           user,
		config:         config,
		authMiddleware: auth,
	}
}

func (s *Route) Initialize() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.response.Success(w, http.StatusOK, "welcome to event planning api!", nil)
	})

	r.HandleFunc("/login", s.user.Login).Methods("POST")
	r.HandleFunc("/register", s.user.Register).Methods("POST")

	api := r.PathPrefix("/api/v1/").Subrouter()

	api.Use(s.authMiddleware.JWTVerify)
	api.HandleFunc("/user", s.user.GetAll).Methods("GET")
	api.HandleFunc("/user/{id}", s.user.GetByID).Methods("GET")
	api.HandleFunc("/user/{id}", s.user.Update).Methods("PUT")
	api.HandleFunc("/user/{id}", s.user.Delete).Methods("DELETE")

	server := http.Server{
		Addr:    s.config.Server.Host + ":" + s.config.Server.Port,
		Handler: r,
	}

	log.Info("Server running!", "PORT", s.config.Server.Port)

	err := server.ListenAndServe()

	if err != nil {
		log.Error(err)
	}
}
