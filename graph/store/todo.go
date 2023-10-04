package store

import (
	"context"
	"go-gql/graph/model"

	"github.com/google/uuid"
)

func (s *store) GetTodos(ctx context.Context) ([]*model.Todo, error) {
	return s.todos, nil
}

func (s *store) CreateTodo(ctx context.Context, in *model.NewTodo) (*model.Todo, error) {
	u, e := s.GetUserByID(ctx, in.UserID)
	if e != nil {
		return nil, e
	}

	newTodo := &model.Todo{
		ID:   uuid.NewString(),
		Text: in.Text,
		Done: false,
		User: u,
	}

	s.todos = append(s.todos, newTodo)

	return newTodo, nil
}
