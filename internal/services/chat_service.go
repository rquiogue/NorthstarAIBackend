package services

import (
	"context"
	"errors"
	"strings"

	"github.com/rquiogue/NorthstarAIBackend/internal/ai"
)

var ErrInvalidMessage = errors.New("message is required")

type ChatService struct {
	client       ai.Client
	defaultModel string
}

func NewChatService(client ai.Client, defaultModel string) *ChatService {
	return &ChatService{client: client, defaultModel: defaultModel}
}

func (s *ChatService) Chat(ctx context.Context, message, model string) (string, error) {
	message = strings.TrimSpace(message)
	if message == "" {
		return "", ErrInvalidMessage
	}
	model = strings.TrimSpace(model)
	if model == "" {
		model = s.defaultModel
	}

	resp, err := s.client.ChatCompletion(ctx, ai.ChatCompletionRequest{
		Message: message,
		Model:   model,
	})
	if err != nil {
		return "", err
	}

	return resp.Response, nil
}
