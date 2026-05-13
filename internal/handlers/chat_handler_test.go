package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

type mockChatService struct {
	response string
	err      error
}

func (m *mockChatService) Chat(_ context.Context, _, _ string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.response, nil
}

func TestChatHandler_Success(t *testing.T) {
	h := NewChatHandler(&mockChatService{response: "hello world"}, zap.NewNop())

	body := bytes.NewBufferString(`{"message":"hi"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/chat", body)
	rec := httptest.NewRecorder()

	h.Chat(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp APIResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if !resp.Success {
		t.Fatal("expected success response")
	}
}

func TestChatHandler_InvalidPayload(t *testing.T) {
	h := NewChatHandler(&mockChatService{}, zap.NewNop())

	req := httptest.NewRequest(http.MethodPost, "/api/v1/chat", bytes.NewBufferString("{"))
	rec := httptest.NewRecorder()

	h.Chat(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}
