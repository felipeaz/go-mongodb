package app

import (
	"github.com/gorilla/mux"
	"go-mongodb/internal/app/users/handler"
)

func Route(userHandler handler.UserHandler) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", userHandler.FindAll).Methods("GET")
	route.HandleFunc("/{id}", userHandler.FindOne).Methods("GET")
	route.HandleFunc("/", userHandler.Create).Methods("POST")
	route.HandleFunc("/{id}", userHandler.UpdateOne).Methods("PUT")
	route.HandleFunc("/{id}", userHandler.Delete).Methods("DELETE")
	return route
}
