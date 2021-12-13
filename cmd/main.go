package main

import (
	"go-mongodb/infra/database/mongodb"
	"go-mongodb/infra/log/logrus"
	"go-mongodb/infra/server/app"
	"go-mongodb/internal/app/users/handler"
	"go-mongodb/internal/app/users/repository"
	"go-mongodb/internal/app/users/service"
	"net/http"
	"os"
)

func main() {
	log := logrus.NewJsonLogger()
	cli, err := mongodb.Connect(os.Getenv("MONGODB_CLUSTER_URI"))
	if err != nil {
		log.Error(err)
	}
	defer func() {
		if err = cli.Disconnect(); err != nil {
			log.Warn(err)
		}
	}()

	collection := cli.C.Database(os.Getenv("sandbox")).Collection("user_sandbox")
	userRepository := repository.NewUserRepository(collection)
	userService := service.NewUserService(userRepository, log)
	userHandler := handler.NewUserHandler(userService, log)

	r := app.Route(userHandler)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
