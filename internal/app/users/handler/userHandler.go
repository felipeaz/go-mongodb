package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go-mongodb/internal/app/domain"
	"io"
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
	users, err := h.UserService.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (h UserHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := h.UserService.FindOne(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	input, err := getInputFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.UserService.Create(input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func (h UserHandler) UpdateOne(w http.ResponseWriter, r *http.Request) {
	input, err := getInputFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	err = h.UserService.UpdateOne(params["id"], input)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, true)
}

func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := h.UserService.Delete(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, true)
}

func getInputFromRequest(r *http.Request) (map[string]interface{}, error) {
	var input map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	return input, nil
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
