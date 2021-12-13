package service

import (
	"encoding/json"
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

func (s UserService) FindAll() ([]domain.User, error) {
	var users []domain.User

	resp, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s UserService) FindOne(id string) (*domain.User, error) {
	var user domain.User

	resp, err := s.UserRepository.FindOne(id)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s UserService) Create(user domain.User) (*domain.User, error) {
	var usr domain.User

	resp, err := s.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &user)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (s UserService) UpdateOne(id string, user domain.User) error {
	return s.UserRepository.UpdateOne(id, user)
}

func (s UserService) Delete(id string) error {
	return s.UserRepository.Delete(id)
}
