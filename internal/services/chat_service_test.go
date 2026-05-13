package services

import (
	"context"
	"errors"
	"testing"

	"github.com/rquiogue/NorthstarAIBackend/internal/ai"
)

type mockAIClient struct {
	response string
	err      error
	gotModel string
}

func (m *mockAIClient) ChatCompletion(_ context.Context, req ai.ChatCompletionRequest) (ai.ChatCompletionResponse, error) {
	m.gotModel = req.Model
	if m.err != nil {
		return ai.ChatCompletionResponse{}, m.err
	}
	return ai.ChatCompletionResponse{Response: m.response}, nil
}

func (m *mockAIClient) StreamChatCompletion(_ context.Context, _ ai.ChatCompletionRequest) (<-chan ai.ChatCompletionChunk, <-chan error) {
	chunks := make(chan ai.ChatCompletionChunk)
	errs := make(chan error)
	close(chunks)
	close(errs)
	return chunks, errs
}

func TestChatService_UsesDefaultModel(t *testing.T) {
	client := &mockAIClient{response: "ok"}
	svc := NewChatService(client, "gpt-default")

	resp, err := svc.Chat(context.Background(), "hello", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp != "ok" {
		t.Fatalf("unexpected response: %s", resp)
	}
	if client.gotModel != "gpt-default" {
		t.Fatalf("expected default model, got %s", client.gotModel)
	}
}

func TestChatService_ValidatesMessage(t *testing.T) {
	svc := NewChatService(&mockAIClient{}, "gpt-default")

	_, err := svc.Chat(context.Background(), "   ", "")
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestChatService_PropagatesClientError(t *testing.T) {
	expectedErr := errors.New("provider failure")
	svc := NewChatService(&mockAIClient{err: expectedErr}, "gpt-default")

	_, err := svc.Chat(context.Background(), "hello", "custom-model")
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected provider error, got %v", err)
	}
}
