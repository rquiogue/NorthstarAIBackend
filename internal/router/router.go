package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rquiogue/NorthstarAIBackend/internal/handlers"
	"github.com/rquiogue/NorthstarAIBackend/internal/middleware"
	"go.uber.org/zap"
)

func New(chatHandler *handlers.ChatHandler, logger *zap.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recovery(logger))
	r.Use(middleware.Logging(logger))
	r.Use(middleware.CORS())

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/chat", chatHandler.Chat)
	})

	return r
}
