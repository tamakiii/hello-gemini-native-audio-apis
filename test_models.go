package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is not set")
	}

	// Based on the pkg.go.dev documentation, the correct usage should be:
	// 1. Create a client
	// 2. Use the client to interact with models

	// Without seeing the exact API, here are a few possibilities:

	// Option 1: Direct model creation
	models := []string{
		"gemini-1.5-flash",
		"gemini-1.5-pro",
		"gemini-2.0-flash-exp",
		"gemini-2.5-flash",
		"gemini-2.5-flash-preview-native-audio-dialog",
	}

	for _, modelName := range models {
		fmt.Printf("\nTesting model: %s\n", modelName)

		// Try different approaches based on common patterns
		// This will need to be adjusted based on the actual API

		// Approach 1: Model might be created directly
		model, err := genai.NewModel(ctx, modelName, genai.WithAPIKey(apiKey))
		if err != nil {
			fmt.Printf("  ❌ Error creating model: %v\n", err)
			continue
		}

		// Approach 2: Or through a client method
		// client := genai.NewClient(...)
		// model := client.Model(modelName)

		// Test generation
		resp, err := model.Generate(ctx, "Say hello")
		if err != nil {
			fmt.Printf("  ❌ Error generating: %v\n", err)
			continue
		}

		fmt.Printf("  ✅ Success! Response: %v\n", resp)
	}
}
