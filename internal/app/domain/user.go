package domain

type UserService interface {
	FindAll() ([]interface{}, error)
	FindOne(id string) (interface{}, error)
	Create(input map[string]interface{}) (interface{}, error)
	UpdateOne(id string, input map[string]interface{}) error
	Delete(id string) error
}

type UserRepository interface {
	FindAll() ([]byte, error)
	FindOne(id string) ([]byte, error)
	Create(input map[string]interface{}) (interface{}, error)
	UpdateOne(id string, input map[string]interface{}) error
	Delete(id string) error
}
