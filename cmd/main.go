package main

import (
	"log"
	"os"
	httpTransport "sessions/internal/http"
	"sessions/internal/service"
	"sessions/internal/storage"

	"github.com/valyala/fasthttp"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://go_user:123@postgres:5432/my_app_db"
	}
	pg, err := storage.New(dsn)

	if err != nil {
		log.Fatal(err)
	}

	sessionService := service.New(pg)
	handler := httpTransport.New(sessionService)

	log.Fatal(fasthttp.ListenAndServe(":8080", httpTransport.Router(handler)))
}
