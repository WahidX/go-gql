package store

import (
	"context"
	"go-gql/todox/model"
)

func (s *store) GetUsers(ctx context.Context) []*model.User {
	return s.users
}
