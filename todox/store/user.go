package store

import (
	"context"
	"go-gql/todox/model"
)

func (s *store) GetUsers(ctx context.Context) []*model.User {
	return s.users
}

func (s *store) CreateUser(ctx context.Context, input *model.User) *model.User {
	s.users = append(s.users, input)
	return input
}
