package app

import (
	"github.com/gorilla/mux"
	"go-mongodb/internal/app/users/handler"
)

func Route(userHandler handler.UserHandler) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/users", userHandler.FindAll).Methods("GET")
	route.HandleFunc("/users/{id}", userHandler.FindOne).Methods("GET")
	route.HandleFunc("/users", userHandler.Create).Methods("POST")
	route.HandleFunc("/users/{id}", userHandler.UpdateOne).Methods("PUT")
	route.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")
	return route
}
