package ai

import (
	"cloud.google.com/go/vertexai/genai"
	"context"
)

// Client defines the interface for AI-related operations
type Client interface {
	Query(ctx context.Context, prompt string, imageURI string) (string, error)
	Close() error
}

// Config holds the configuration for the AI client
type Config struct {
	ProjectID string
	Region    string
	ModelName string
}

// vertexAIClient is the AI client implementation
type vertexAIClient struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

// NewClient creates a new AI client
func NewClient(ctx context.Context, config *Config) (Client, error) {
	client, err := genai.NewClient(ctx, config.ProjectID, config.Region)
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel(config.ModelName)

	return &vertexAIClient{
		client: client,
		model:  model,
	}, nil
}

func (v *vertexAIClient) Query(ctx context.Context, prompt string, imageURI string) (string, error) {
	var parts []genai.Part
	parts = append(parts, genai.Text(prompt))

	if imageURI != "" {
		img := genai.FileData{
			MIMEType: "image/jpeg",
			FileURI:  imageURI,
		}
		parts = append(parts, img)
	}

	// TODO: GenerateContentStream()を使うかどうか検討
	res, err := v.model.GenerateContent(ctx, parts...)
	if err != nil {
		return "", err
	}

	// Assuming we want the text from the first candidate's first part...?
	if len(res.Candidates) > 0 && len(res.Candidates[0].Content.Parts) > 0 {
		if text, ok := res.Candidates[0].Content.Parts[0].(genai.Text); ok {
			return string(text), nil
		}
	}

	return "", nil
}

func (v *vertexAIClient) Close() error {
	return v.client.Close()
}
