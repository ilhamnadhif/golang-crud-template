package repository

import (
	"context"
	"golang-crud-template/db/entity"
)

type PostRepository interface {
	Insert(ctx context.Context, post entity.Post) (entity.Post, error)
	FindById(ctx context.Context, id int32) (entity.Post, error)
	FindAll(ctx context.Context) ([]entity.Post, error)
	UpdateById(ctx context.Context, id int32, comment entity.Post) (entity.Post, error)
	DeleteById(ctx context.Context, id int32) (string, error)
}
