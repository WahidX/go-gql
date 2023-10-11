package store

import (
	"context"
	"go-gql/todox/model"
)

type Store interface {
	GetUsers(ctx context.Context) []*model.User
	GetTodos(ctx context.Context) []*model.Todo
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
