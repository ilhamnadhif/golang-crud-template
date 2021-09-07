package main

import (
	"context"
	"embed"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"golang-crud-template/db"
	"golang-crud-template/db/entity"
	"golang-crud-template/db/repository"
	"golang-crud-template/helper"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:embed view/*.gohtml
var templates embed.FS

var MyTemplate = template.Must(template.New("").Funcs(map[string]interface{}{
	"parsedate": func(value time.Time) string {
		year, month, day := value.Date()
		return strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year)
	},
}).ParseFS(templates, "view/*.gohtml"))

var Home = func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	MyTemplate.ExecuteTemplate(writer, "index", nil)
}

var Blog = func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()

	posts, err := postRepository.FindAll(ctx)
	helper.ErrorHandling(err)

	MyTemplate.ExecuteTemplate(writer, "blog", map[string]interface{}{
		"Posts": posts,
	})
}

var CreateBlog = func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	title := request.PostFormValue("title")
	body := request.PostFormValue("body")
	// image
	file, fileHeader, err := request.FormFile("image")
	helper.ErrorHandling(err)
	imgName := strings.Join(strings.Split(fileHeader.Filename, " "), "_")
	fileDestination, err := os.Create("./images/" + imgName)
	helper.ErrorHandling(err)
	_, err = io.Copy(fileDestination, file)
	helper.ErrorHandling(err)

	slug := strings.Join(strings.Split(strings.ToLower(title), " "), "-")

	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	post := entity.Post{
		Title: title,
		Slug:  slug,
		Body:  body,
		Image: imgName,
	}
	result, err := postRepository.Insert(ctx, post)
	helper.ErrorHandling(err)
	fmt.Println(result)
	http.Redirect(writer, request, "/blog", 301)
}

var DetailBlog = func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	slug := params.ByName("slug")
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	post, err := postRepository.FindBySlug(ctx, slug)
	helper.ErrorHandling(err)
	MyTemplate.ExecuteTemplate(writer, "detail", post)
}

var DeleteBlog = func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	blogId := params.ByName("id")
	id, _ := strconv.Atoi(blogId)
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	post, err := postRepository.FindById(ctx, int32(id))
	helper.ErrorHandling(err)
	_, error := postRepository.DeleteById(ctx, post.Id)
	helper.ErrorHandling(error)
	RemoveImage(post.Image)
	http.Redirect(writer, request, "/blog", 301)
}

var FormEdit = func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	blogId := params.ByName("id")
	id, _ := strconv.Atoi(blogId)
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	post, err := postRepository.FindById(ctx, int32(id))
	helper.ErrorHandling(err)
	MyTemplate.ExecuteTemplate(writer, "formEdit", map[string]interface{}{
		"Post": post,
	})
}

var UpdateBlogById = func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := request.PostFormValue("id")
	idBlog, _ := strconv.Atoi(id)
	title := request.PostFormValue("title")
	body := request.PostFormValue("body")
	slug := strings.Join(strings.Split(strings.ToLower(title), " "), "-")

	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()
	result, err := postRepository.FindById(ctx, int32(idBlog))
	helper.ErrorHandling(err)
	var post entity.Post = entity.Post{
		Title: title,
		Slug:  slug,
		Body:  body,
	}
	_, error := postRepository.UpdateById(ctx, result.Id, post)
	helper.ErrorHandling(error)
	http.Redirect(writer, request, "/blog", 301)
}

////go:embed images
//var images embed.FS

func main() {
	router := httprouter.New()
	router.ServeFiles("/images/*filepath", http.Dir("./images/"))
	router.GET("/", Home)
	router.GET("/blog", Blog)
	router.GET("/blog/:slug", DetailBlog)
	router.GET("/edit/:id", FormEdit)

	router.POST("/create", CreateBlog)
	router.GET("/delete/:id", DeleteBlog)
	router.POST("/update", UpdateBlogById)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.ErrorHandling(err)
}

func RemoveImage(image string) {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	directory := filepath.Join(d, "./images")
	e := os.Remove(directory + "/" + image)
	if e != nil {
		panic(e)
	}
}
