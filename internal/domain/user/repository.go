package user

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id int64) (*User, error)
	Update(user *User) error
	Delete(id int64) error
}
