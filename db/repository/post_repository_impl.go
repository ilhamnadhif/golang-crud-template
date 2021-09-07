package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud-template/db/entity"
	"strconv"
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
	script := "SELECT id, title, slug, body, image, created_at FROM posts WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	var post entity.Post
	if err != nil {
		return post, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&post.Id, &post.Title, &post.Slug, &post.Body, &post.Image, &post.CreateAt)
		return post, nil
	} else {
		// nggak ada
		return post, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}
func (repository *postRepositoryImpl) FindBySlug(ctx context.Context, slug string) (entity.Post, error) {
	script := "SELECT id, title, slug, body, image, created_at FROM posts WHERE slug = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, slug)
	var post entity.Post
	if err != nil {
		return post, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&post.Id, &post.Title, &post.Slug, &post.Body, &post.Image, &post.CreateAt)
		return post, nil
	} else {
		// nggak ada
		return post, errors.New(("Slug " + slug + " Not Found"))
	}
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
	script := "DELETE FROM posts WHERE id = ?"
	_, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return "Id " + strconv.Itoa(int(id)) + " Gagal Dihapus", err
	}
	return "Id " + strconv.Itoa(int(id)) + " Berhasil Dihapus", nil
}
