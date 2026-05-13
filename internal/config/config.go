package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const defaultAIBaseURL = "https://api.openai.com"

type Config struct {
	Port           string
	AIAPIKey       string
	AIBaseURL      string
	DefaultModel   string
	AllowedOrigins []string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:           strings.TrimSpace(os.Getenv("PORT")),
		AIAPIKey:       strings.TrimSpace(os.Getenv("AI_API_KEY")),
		AIBaseURL:      strings.TrimSpace(os.Getenv("AI_BASE_URL")),
		DefaultModel:   strings.TrimSpace(os.Getenv("AI_DEFAULT_MODEL")),
		AllowedOrigins: parseOrigins(strings.TrimSpace(os.Getenv("CORS_ALLOWED_ORIGINS"))),
	}

	if cfg.AIBaseURL == "" {
		cfg.AIBaseURL = defaultAIBaseURL
	}
	if cfg.DefaultModel == "" {
		cfg.DefaultModel = "gpt-4o-mini"
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	var missing []string
	if c.AIAPIKey == "" {
		missing = append(missing, "AI_API_KEY")
	}
	if c.Port == "" {
		missing = append(missing, "PORT")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required env vars: %s", strings.Join(missing, ", "))
	}

	return nil
}

func parseOrigins(raw string) []string {
	if raw == "" {
		return []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	}

	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))
	for _, part := range parts {
		origin := strings.TrimSpace(part)
		if origin != "" {
			origins = append(origins, origin)
		}
	}
	if len(origins) == 0 {
		return []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	}

	return origins
}
