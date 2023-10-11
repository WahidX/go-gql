package store

import (
	"context"
	"go-gql/todox/model"
)

type Store interface {
	// Users
	CreateUser(ctx context.Context, input *model.User) *model.User
	GetUsers(ctx context.Context) []*model.User

	CreateTodo(ctx context.Context, input *model.Todo) *model.Todo
	GetTodos(ctx context.Context) []*model.Todo
	UserTodos(ctx context.Context, userID string) []*model.Todo
}

type store struct {
	users []*model.User
	todos []*model.Todo
}

func New() Store {
	return &store{
		users: []*model.User{},
		todos: []*model.Todo{},
	}
}
