package domain

type User struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type UserService interface {
}

type UserRepository interface {
}
