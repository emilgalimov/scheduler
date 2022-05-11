package app

import pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"

type tserver struct {
	pb.UnimplementedSomeStuffServer

	repo Repository
}

func NewServer(repo Repository) *tserver {
	return &tserver{repo: repo}
}
