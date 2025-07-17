package command_test

import (
	"patterns/behavioral/command"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Command(t *testing.T) {
	t.Parallel()
	// Arrange
	cmd := &command.CreateUserCommand{Name: "John Doe", Email: "john@gmail.com"}
	handler := command.NewCreateUserHandler(command.NewUserRepository())

	// Act
	err := handler.Handle(cmd)

	// Assert
	require.NoError(t, err)
}

func Test_InvalidCommand(t *testing.T) {
	t.Parallel()
	handler := command.NewCreateUserHandler(command.NewUserRepository())

	err := handler.Handle("invalid command")

	require.Equal(t, "invalid command", err.Error())
}
