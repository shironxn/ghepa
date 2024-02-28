package route

import (
	"ghepa/config"
	"ghepa/internal/core/port"
	"ghepa/internal/middleware"
	"ghepa/internal/util"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
)

type Route struct {
	user           port.UserHandler
	event          port.EventHandler
	comment        port.CommentHandler
	response       util.Response
	authMiddleware middleware.AuthMiddleware
	config         *config.App
}

func NewRoute(user port.UserHandler, event port.EventHandler, comment port.CommentHandler, auth middleware.AuthMiddleware, config *config.App) *Route {
	return &Route{
		user:           user,
		event:          event,
		comment:        comment,
		config:         config,
		authMiddleware: auth,
	}
}

func (s *Route) Initialize() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.response.Success(w, http.StatusOK, "welcome to event planning api", nil)
	})

	auth := r.PathPrefix("/api/v1/auth/").Subrouter()
	auth.HandleFunc("/login", s.user.Login).Methods("POST")
	auth.HandleFunc("/register", s.user.Register).Methods("POST")

	api := r.PathPrefix("/api/v1/").Subrouter()
	api.Use(s.authMiddleware.JWTVerify)

	api.HandleFunc("/user", s.user.GetAll).Methods("GET")
	api.HandleFunc("/user/{id}", s.user.GetByID).Methods("GET")
	api.HandleFunc("/user/{id}", s.user.Update).Methods("PUT")
	api.HandleFunc("/user/{id}", s.user.Delete).Methods("DELETE")

	api.HandleFunc("/event", s.event.Create).Methods("POST")
	api.HandleFunc("/event", s.event.GetAllByUser).Methods("GET")
	api.HandleFunc("/event/{id}", s.event.GetByID).Methods("GET")
	api.HandleFunc("/event/{id}", s.event.Update).Methods("PUT")
	api.HandleFunc("/event/{id}", s.event.Delete).Methods("DELETE")
	api.HandleFunc("/event/{id}/join", s.event.JoinEvent).Methods("POST")

	api.HandleFunc("/comment/{id}", s.comment.Create).Methods("POST")
	api.HandleFunc("/comment", s.comment.GetAll).Methods("GET")

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
