package ai

import (
	"context"

	"github.com/anuragdaksh7/zapmail-backend/logger"
	"google.golang.org/genai"
)

func ChatResultToString(result *genai.GenerateContentResponse) string {
	var output string
	for _, candidate := range result.Candidates {
		for _, part := range candidate.Content.Parts {
			output += part.Text // Direct access
		}
	}
	return output
}

func GenerateGemmaResponse(ctx context.Context, input string) (string, error) {
	chat, err := GeminiClient.Chats.Create(ctx, "gemma-3-1b-it", nil, nil)
	if err != nil {
		return "", err
	}

	logger.Logger.Info("Input: " + input)
	result, err := chat.SendMessage(ctx, genai.Part{
		Text: input,
	})
	text := ChatResultToString(result)
	logger.Logger.Info("Output + " + text)

	return text, err
}
