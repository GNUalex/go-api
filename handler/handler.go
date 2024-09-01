package handler

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	PublishedAt string `json:"publishedAt"`
	Author      string `json:"author"`
	Content     string `json:"content"`
}

func GetPosts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var posts []Post
		db.Find(&posts)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(posts)
	}
}

func GetPost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post *Post
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err := db.Where("slug = ?", r.PathValue("slug")).First(&post).Error; err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Not found")
		} else {
			json.NewEncoder(w).Encode(post)
		}
	}
}
