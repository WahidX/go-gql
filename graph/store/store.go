package store

import (
	"context"
	"go-gql/graph/model"
)

type store struct {
	todos []*model.Todo
	users []*model.User
}

type Store interface {
	GetTodos(ctx context.Context) ([]*model.Todo, error)
	CreateTodo(ctx context.Context, in *model.NewTodo) (*model.Todo, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

func New() Store {
	return &store{
		todos: []*model.Todo{},
		users: []*model.User{
			{ID: "123", Name: "Goku"},
			{ID: "124", Name: "Gohan"},
			{ID: "125", Name: "Picolo"},
		},
	}
}
