package repository

import (
	"context"
	"database/sql"
	"golang-crud-template/db/entity"
)

type postRepositoryImpl struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepositoryImpl{DB: db}
}

func (repository *postRepositoryImpl) Insert(ctx context.Context, post entity.Post) (entity.Post, error) {
	script := "INSERT INTO posts(title, slug, body, image) VALUES (?, ?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, post.Title, post.Slug, post.Body, post.Image)
	if err != nil {
		return post, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return post, err
	}
	post.Id = int32(id)
	return post, nil
}

func (repository *postRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Post, error) {
	panic("asd")
	//script := "SELECT id, title, slug, body, image, created_at FROM posts"
	//rows, err := repository.DB.QueryContext(ctx, script)
	//if err != nil {
	//	return pos, err
	//}
}

func (repository *postRepositoryImpl) FindAll(ctx context.Context) ([]entity.Post, error) {

	script := "SELECT id, title, slug, body, image, created_at FROM posts ORDER BY created_at DESC "
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		rows.Scan(&post.Id, &post.Title, &post.Slug, &post.Body, &post.Image, &post.CreateAt)
		posts = append(posts, post)
	}
	return posts, nil
}

func (repository *postRepositoryImpl) UpdateById(ctx context.Context, id int32, comment entity.Post) (entity.Post, error) {
	panic("implement me")
}

func (repository *postRepositoryImpl) DeleteById(ctx context.Context, id int32) (string, error) {
	panic("implement me")
}
