package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	PublishedAt string `json:"publishedAt"`
	Author      string `json:"author"`
	Content     string `json:"content"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts = []Post{
		{
			Id:          1,
			Title:       "Sample blog post 1",
			Slug:        "sample-blog-post-1",
			PublishedAt: time.Now().String(),
			Author:      "GNUalex",
			Content:     "Proin tristique sagittis nisl, sit amet ultricies nulla vehicula nec. Ut imperdiet lacus eu porttitor feugiat. Aenean ipsum risus, rhoncus ac tempor at, rhoncus vitae dui. Aliquam erat volutpat. Morbi non vulputate ipsum.",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(r.PathValue("slug"))
}
