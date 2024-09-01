package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string `json:"title"`
	Slug        string `gorm:"unique" json:"slug"`
	PublishedAt string `json:"publishedAt"`
	Author      string `json:"author"`
	Content     string `json:"content"`
}

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Post{})
	return db
}
