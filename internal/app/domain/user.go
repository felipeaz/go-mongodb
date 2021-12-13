package domain

type User struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type UserService interface {
	FindAll() ([]User, error)
	FindOne(id string) (*User, error)
	Create(user User) (*User, error)
	UpdateOne(id string, user User) error
	Delete(id string) error
}

type UserRepository interface {
	FindAll() ([]byte, error)
	FindOne(id string) ([]byte, error)
	Create(user User) ([]byte, error)
	UpdateOne(id string, user User) error
	Delete(id string) error
}
