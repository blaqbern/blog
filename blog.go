package main

import (
	"encoding/json"
	"fmt"
	"github.com/blaqbern/blog/cors"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	Title string
	Body  []byte
}

func (p *Post) save() error {
	filename := p.Title + ".md"
	return ioutil.WriteFile(filename, p.Body, 0600)
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
	filename := title + ".md"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Post{Title: title, Body: body}, nil
}

func getPostListHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, _ := getPostList() // @todo handle the error
	json.NewEncoder(w).Encode(data)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/post/"):]
	w.Header().Set("Access-Control-Allow-Origin", "*")
	p, _ := getPost(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/posts", getPostListHandler)
	http.HandleFunc("/post/", getPostHandler)
	log.Fatal(http.ListenAndServe(":5555", nil))
}
