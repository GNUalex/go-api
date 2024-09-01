package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GNUalex/go-api/handler"
	"github.com/GNUalex/go-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HttpServer struct {
	DB *gorm.DB
}

func (httpServer *HttpServer) Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_TIMEZONE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database")
	}

	httpServer.DB = model.Migrate(db)
}

func (httpServer *HttpServer) ListenAndServe(port string) {
	httpServer.Init()
	http.HandleFunc("/posts", handler.GetPosts(httpServer.DB))
	http.HandleFunc("/post/{slug}", handler.GetPost(httpServer.DB))

	log.Fatal(http.ListenAndServe(port, nil))
}
