package transport

import (
	"net/http"

	"github.com/adedaramola/golang-auth/services"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserService *services.UserService
}

func NewHandler(userService *services.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) RegisterRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	}).Methods("GET")

	router.HandleFunc("/register", h.registerUser).Methods("POST")

	return router
}
