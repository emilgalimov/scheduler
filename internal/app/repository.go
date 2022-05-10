package app

import "context"

type Repository interface {
	CreateUser(context.Context) (uint64, error)
}
