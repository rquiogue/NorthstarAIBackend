package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const defaultAIBaseURL = "https://api.openai.com"

type Config struct {
	Port         string
	AIAPIKey     string
	AIBaseURL    string
	DefaultModel string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:         strings.TrimSpace(os.Getenv("PORT")),
		AIAPIKey:     strings.TrimSpace(os.Getenv("AI_API_KEY")),
		AIBaseURL:    strings.TrimSpace(os.Getenv("AI_BASE_URL")),
		DefaultModel: strings.TrimSpace(os.Getenv("AI_DEFAULT_MODEL")),
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

	if c.AIBaseURL == "" {
		return errors.New("AI_BASE_URL cannot be empty when provided")
	}

	return nil
}
