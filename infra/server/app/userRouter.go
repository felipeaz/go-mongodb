package app

import (
	"github.com/gorilla/mux"
	"go-mongodb/internal/app/users/handler"
)

func Route(userHandler handler.UserHandler) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", userHandler.FindAll)
	route.HandleFunc("/{id}", userHandler.FindOne)
	route.HandleFunc("/", userHandler.Create)
	route.HandleFunc("/{id}", userHandler.UpdateOne)
	route.HandleFunc("/{id}", userHandler.Delete)
	return route
}
