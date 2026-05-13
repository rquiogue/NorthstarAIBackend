package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rquiogue/NorthstarAIBackend/internal/services"
	"go.uber.org/zap"
)

type ChatService interface {
	Chat(ctx context.Context, message, model string) (string, error)
}

type ChatHandler struct {
	service ChatService
	logger  *zap.Logger
}

type ChatRequest struct {
	Message string `json:"message"`
	Model   string `json:"model"`
}

type ChatResponseData struct {
	Response string `json:"response"`
}

func NewChatHandler(service ChatService, logger *zap.Logger) *ChatHandler {
	return &ChatHandler{service: service, logger: logger}
}

func (h *ChatHandler) Chat(w http.ResponseWriter, r *http.Request) {
	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	resp, err := h.service.Chat(r.Context(), req.Message, req.Model)
	if err != nil {
		if errors.Is(err, services.ErrInvalidMessage) {
			WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		h.logger.Error("chat request failed", zap.Error(err))
		WriteError(w, http.StatusBadGateway, "failed to process AI request")
		return
	}

	WriteSuccess(w, http.StatusOK, ChatResponseData{Response: resp})
}
