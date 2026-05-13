# NorthstarAIBackend

A minimal Go HTTP server with one endpoint to build on.

## Run

```bash
go run main.go
```

## Endpoint

`POST /chat`

Request:
```json
{ "message": "hello" }
```

Response:
```json
{ "reply": "echo: hello" }
```
