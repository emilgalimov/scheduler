package app

import (
	"context"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *tserver) UnsubscribeUser(ctx context.Context, req *pb.UnsubscribeUserRequest) (*pb.UnsubscribeUserResponse, error) {

	//TODO добавить проверку наличия пользователя

	if _, err := t.repo.GetTask(ctx, req.TaskID); err != nil {
		return nil, status.Errorf(codes.NotFound, "task with ID %v not found", req.TaskID)
	}

	if _, err := t.repo.GetUserTask(ctx, req.UserID, req.TaskID); err != nil {
		return nil, status.Errorf(codes.NotFound, "user not subscribed")
	}

	err := t.repo.DeleteUserTask(ctx, req.UserID, req.TaskID)
	if err != nil {
		return nil, err
	}
	return &pb.UnsubscribeUserResponse{IsUnsubscribed: true}, nil
}
