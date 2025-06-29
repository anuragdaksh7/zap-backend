package ai

import (
	"context"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"google.golang.org/genai"
)

var GeminiClient *genai.Client

func InitGemini() error {
	_config, err := config.LoadConfig(".")
	if err != nil {
		return err
	}

	ctx := context.Background()

	cfg := &genai.ClientConfig{
		APIKey: _config.GeminiAPIKey,
	}

	c, err := genai.NewClient(ctx, cfg)
	if err != nil {
		return err
	}
	GeminiClient = c

	return nil
}
