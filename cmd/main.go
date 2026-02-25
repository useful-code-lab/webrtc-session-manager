package main

import (
	"log"
	httpTransport "sessions/internal/http"
	"sessions/internal/service"
	"sessions/internal/storage"

	"github.com/valyala/fasthttp"
)

func main() {
	pg, err := storage.New("postgres://go_user:123@postgres:5432/my_app_db")

	if err != nil {
		log.Fatal(err)
	}

	sessionService := service.New(pg)
	handler := httpTransport.New(sessionService)

	log.Fatal(fasthttp.ListenAndServe(":8080", httpTransport.Router(handler)))
}
