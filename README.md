# NorthstarAIBackend

Stateless Go AI API service scaffold.

## Run

1. Copy `.env.example` to `.env` and set `AI_API_KEY`
2. Start the API:

```bash
make run
```

## API

- `POST /api/v1/chat`

Request:

```json
{
  "message": "hello",
  "model": "gpt-4o-mini"
}
```

Response:

```json
{
  "success": true,
  "data": {
    "response": "..."
  },
  "error": null
}
```
