// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Authors, error)
	DeleteAuthor(ctx context.Context, id int64) error
	GetAuthor(ctx context.Context, id int64) (Authors, error)
	ListAuthors(ctx context.Context) ([]Authors, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error
}

var _ Querier = (*Queries)(nil)
