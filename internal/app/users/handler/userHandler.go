package handler

import (
	"github.com/sirupsen/logrus"
	"go-mongodb/internal/app/domain"
	"net/http"
)

type UserHandler struct {
	UserService domain.UserService
	Log         *logrus.Logger
}

func NewUserHandler(service domain.UserService, log *logrus.Logger) UserHandler {
	return UserHandler{
		UserService: service,
		Log:         log,
	}
}

func (h UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) FindOne(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) UpdateOne(w http.ResponseWriter, r *http.Request) {

}

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
