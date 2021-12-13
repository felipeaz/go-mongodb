package main

import (
	"go-mongodb/infra/database/mongodb"
	"go-mongodb/infra/log/logrus"
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

}
