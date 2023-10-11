package store

import (
	"context"
	"go-gql/todox/model"
)

func (s *store) GetTodos(ctx context.Context) []*model.Todo {
	return s.todos
}

func (s *store) CreateTodo(ctx context.Context, input *model.Todo) *model.Todo {
	s.todos = append(s.todos, input)
	return input
}

func (s *store) UserTodos(ctx context.Context, userID string) []*model.Todo {
	t := []*model.Todo{}
	for _, todo := range s.todos {
		if todo.User.ID == userID {
			t = append(t, todo)
		}
	}

	return t
}
