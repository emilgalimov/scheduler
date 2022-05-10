package app

import (
	"context"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
)

func (t *tserver) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserID, error) {
	ID, err := t.repo.CreateUser(ctx)

	return &pb.UserID{ID: ID}, err
}
