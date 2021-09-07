package test

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-crud-template/db"
	"golang-crud-template/db/entity"
	"golang-crud-template/db/repository"
	"golang-crud-template/helper"
	"path"
	"runtime"
	"testing"
)

func TestPostInsert(t *testing.T) {
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	post := entity.Post{
		Title: "Ini Title 1",
		Slug:  "ini-title-1",
		Body:  "Ini deskripsi body 1",
		Image: "image1.png",
	}
	result, err := postRepository.Insert(ctx, post)
	helper.ErrorHandling(err)
	fmt.Println(result)
}

func TestPostFindAll(t *testing.T) {
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	posts, err := postRepository.FindAll(ctx)
	helper.ErrorHandling(err)
	for _, post := range posts {
		fmt.Println(post)
	}
}
func TestPostFindById(t *testing.T) {
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	posts, err := postRepository.FindById(ctx, 9)
	helper.ErrorHandling(err)
	fmt.Println(posts)
}
func TestPostFindBySlug(t *testing.T) {
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	posts, err := postRepository.FindBySlug(ctx, "ini-title-1")
	helper.ErrorHandling(err)
	fmt.Println(posts)
}
func TestDeleteById(t *testing.T) {
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	posts, err := postRepository.DeleteById(ctx, 3)
	helper.ErrorHandling(err)
	fmt.Println(posts)
}

func TestGetPwd(t *testing.T) {
	_, b, _ , _ := runtime.Caller(0)
	fmt.Println(b)
	d := path.Join(path.Dir(b))
	fmt.Println(d)
}