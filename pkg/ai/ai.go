package ai

import (
	"context"
	"errors"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/vertexai/genai"
)

// Client defines the interface for AI-related operations
type Client interface {
	Query(ctx context.Context, prompt string, imageURI string) (<-chan string, <-chan error)
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

// Query sends a query to the AI model
func (v *vertexAIClient) Query(ctx context.Context, prompt string, imageURI string) (<-chan string, <-chan error) {
	responseChan := make(chan string)
	errorChan := make(chan error, 1)

	go func() {
		defer close(responseChan)
		defer close(errorChan)

		var parts []genai.Part
		parts = append(parts, genai.Text(prompt))

		if imageURI != "" {
			img := genai.FileData{
				MIMEType: "image/jpeg",
				FileURI:  imageURI,
			}
			parts = append(parts, img)
		}

		iter := v.model.GenerateContentStream(ctx, parts...)
		for {
			resp, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				return
			}
			if err != nil {
				errorChan <- err
				return
			}

			if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
				if text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
					responseChan <- string(text)
				}
			}
		}
	}()

	return responseChan, errorChan
}

// Close closes the AI client
func (v *vertexAIClient) Close() error {
	return v.client.Close()
}
