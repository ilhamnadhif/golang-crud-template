package main

import (
	"context"
	"embed"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-crud-template/db"
	"golang-crud-template/db/entity"
	"golang-crud-template/db/repository"
	"golang-crud-template/helper"
	"html/template"
	"io"
	"net/http"
	"os"
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

var Home = func(writer http.ResponseWriter, request *http.Request) {
	MyTemplate.ExecuteTemplate(writer, "index", nil)
}

var Blog = func(writer http.ResponseWriter, request *http.Request) {
	postRepository := repository.NewPostRepository(db.GetConnection())
	ctx := context.Background()

	posts, err := postRepository.FindAll(ctx)
	helper.ErrorHandling(err)

	MyTemplate.ExecuteTemplate(writer, "blog", map[string]interface{}{
		"Posts": posts,
	})
}

var CreateBlog = func(writer http.ResponseWriter, request *http.Request) {
	defer http.Redirect(writer, request, "/blog", http.StatusTemporaryRedirect)
	title := request.PostFormValue("title")
	body := request.PostFormValue("body")
	// image
	file, fileHeader, err := request.FormFile("image")
	helper.ErrorHandling(err)
	fileDestination, err := os.Create("./images/" + fileHeader.Filename)
	helper.ErrorHandling(err)
	_, err = io.Copy(fileDestination, file)
	helper.ErrorHandling(err)
	imgName := fileHeader.Filename

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
}
var DetailBlog = func(writer http.ResponseWriter, request *http.Request) {
	MyTemplate.ExecuteTemplate(writer, "detail", nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/blog/", Blog)
	mux.HandleFunc("/blog/create", CreateBlog)
	//mux.HandleFunc("/blog/?slug=", DetailBlog)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	helper.ErrorHandling(err)
}
