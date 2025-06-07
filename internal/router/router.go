package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"myapp/internal/handler"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", handler.Ping)
	return r
}
