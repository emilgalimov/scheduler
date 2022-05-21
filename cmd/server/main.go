package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/app"
	"gitlab.ozon.dev/emilgalimov/homework-2/internal/repository"
	pb "gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	cfg, err := NewConfig("/app/config.yaml")
	if err != nil {
		panic(fmt.Sprint("error load config ", err))
	}

	ctx := context.Background()

	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DbName,
	)

	conn, _ := pgxpool.Connect(ctx, connectString)

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("error pinging db: ", err)
	}

	repo := repository.NewTasksRepository(conn)

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := app.NewServer(repo)
	grpcServer := grpc.NewServer()

	pb.RegisterSmartCalendarServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
