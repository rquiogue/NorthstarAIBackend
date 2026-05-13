package middleware

import (
	"fmt"
	"net/http"

	"github.com/rquiogue/NorthstarAIBackend/internal/handlers"
	"go.uber.org/zap"
)

func Recovery(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					logger.Error("panic recovered",
						zap.String("request_id", GetRequestID(r.Context())),
						zap.String("panic", fmt.Sprint(rec)),
					)
					handlers.WriteError(w, http.StatusInternalServerError, "internal server error")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
