package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/googleapis/go-genai"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is not set")
	}

	// Create a client with the correct API
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGoogleAI,
	})
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	// Test different model names including the native audio dialog model
	models := []string{
		"gemini-1.5-flash",
		"gemini-1.5-pro",
		"gemini-2.0-flash-exp",
		"gemini-2.5-flash",
		"gemini-2.5-flash-preview-native-audio-dialog",
	}

	for _, modelName := range models {
		fmt.Printf("\nTesting model: %s\n", modelName)

		model := client.GenerativeModel(modelName)

		// Simple test prompt
		resp, err := model.GenerateContent(ctx, genai.Text("Say hello"))
		if err != nil {
			fmt.Printf("  ❌ Error: %v\n", err)
		} else {
			fmt.Printf("  ✅ Success!\n")
			for _, candidate := range resp.Candidates {
				if candidate.Content != nil {
					for _, part := range candidate.Content.Parts {
						fmt.Printf("     Response: %v\n", part)
					}
				}
			}
		}
	}
}
