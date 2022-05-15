package repository

import "context"

type Connection interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) Row
}

type Row interface {
	Scan(dest ...interface{}) error
}
