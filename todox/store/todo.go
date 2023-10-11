package store

import (
	"context"
	"go-gql/todox/model"
)

func (s *store) GetTodos(ctx context.Context) []*model.Todo {
	return s.todos
}
