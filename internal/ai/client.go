package ai

import "context"

type ChatCompletionRequest struct {
	Message string
	Model   string
	Stream  bool
}

type ChatCompletionResponse struct {
	Response string
}

type ChatCompletionChunk struct {
	Delta string
	Done  bool
}

type Client interface {
	ChatCompletion(ctx context.Context, req ChatCompletionRequest) (ChatCompletionResponse, error)
	StreamChatCompletion(ctx context.Context, req ChatCompletionRequest) (<-chan ChatCompletionChunk, <-chan error)
}
