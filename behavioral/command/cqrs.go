package command

import "errors"

type Command any

type CreateUserCommand struct {
	Name  string
	Email string
}

type CommandHandler interface {
	Handle(cmd Command) error
}

type createUserHandler struct {
	repo UserRepository
}

func NewCreateUserHandler(repo UserRepository) CommandHandler {
	return createUserHandler{repo: repo}
}

func (h createUserHandler) Handle(cmd Command) error {
	c, ok := cmd.(*CreateUserCommand)
	if !ok {
		return errors.New("invalid command")
	}

	return h.repo.Save(User{Name: c.Name, Email: c.Email})
}
