package test

import (
	"context"
	"fmt"
	"golang-crud-template/db"
	"golang-crud-template/db/entity"
	"golang-crud-template/db/repository"
	"golang-crud-template/helper"
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
