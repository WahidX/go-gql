package store

import (
	"context"
	"errors"
	"go-gql/graph/model"
)

func (s *store) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, errors.New("User not found")
}
