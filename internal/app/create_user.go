package app

import (
	"context"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (t *tserver) CreateUser(ctx context.Context, _ *pb.CreateUserRequest) (*pb.User, error) {
	ID, err := t.repo.CreateUser(ctx)

	return &pb.User{ID: ID}, err
}
