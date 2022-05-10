package app

import pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"

type tserver struct {
	repo Repository
	pb.UnimplementedSomeStuffServer
}

func NewServer(repo Repository) *tserver {
	return &tserver{repo: repo}
}
