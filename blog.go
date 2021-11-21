package main

import (
	"encoding/json"
	"github.com/blaqbern/blog/internal/middleware"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func getPostList() ([]string, error) {
	files, err := ioutil.ReadDir("./posts")
	if err != nil {
		return nil, err
	}

	filenames := make([]string, len(files))
	for i, f := range files {
		filenames[i] = f.Name()
	}

	return filenames, nil
}

func getPost(title string) (*Post, error) {
	filename := "./posts/" + title + ".md"
	body, err := ioutil.ReadFile(filename)
	log.Printf("title = %v; body = %v", title, string(body))
	if err != nil {
		log.Printf("err = %v", err)
		return nil, err
	}
	return &Post{Title: title, Body: string(body)}, nil
}

func getPostListHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := getPostList() // @todo handle the error
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/posts/"):]
	p, _ := getPost(title)

	log.Printf("post = %v", p)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func main() {
	http.HandleFunc("/posts", middleware.WithMiddleware(getPostListHandler))
	http.HandleFunc("/posts/", middleware.WithMiddleware(getPostHandler))
	log.Fatal(http.ListenAndServe(":5555", nil))
}
