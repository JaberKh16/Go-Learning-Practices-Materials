package domain

type UserRepository interface {
	Save(user User) error
	GetAll() []User
	GetByID(id int) (User, error)
	FindByEmail(email string) (User, error)
	Update(user User) error
	Delete(id int) error
}