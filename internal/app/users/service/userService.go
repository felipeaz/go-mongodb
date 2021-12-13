package service

import (
	"github.com/sirupsen/logrus"
	"go-mongodb/internal/app/domain"
)

type UserService struct {
	UserRepository domain.UserRepository
	Log            *logrus.Logger
}

func NewUserService(repository domain.UserRepository, log *logrus.Logger) UserService {
	return UserService{
		UserRepository: repository,
		Log:            log,
	}
}
