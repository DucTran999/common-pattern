package command

import "fmt"

type UserRepository interface {
	Save(user User) error
}

type userRepo struct {
}

func NewUserRepository() *userRepo {
	return &userRepo{}
}

func (r *userRepo) Save(u User) error {
	fmt.Println("Saving user:", u.Name, "with email:", u.Email)

	return nil
}
