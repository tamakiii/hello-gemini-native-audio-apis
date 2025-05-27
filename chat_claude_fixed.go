package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

var model = flag.String("model", "gemini-2.0-flash", "the model name, e.g. models/gemini-2.0-flash")
var listModels = flag.Bool("list-models", false, "list available models")

func chat(ctx context.Context, apiKey string) {
	// Create a client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	if *listModels {
		fmt.Println("Listing available models...")
		fmt.Println("Available models:")
		
		for model, err := range client.Models.All(ctx) {
			if err != nil {
				log.Fatalf("Failed to get model: %v", err)
			}
			fmt.Printf("- Name: %s\n", model.Name)
			fmt.Printf("  Version: %s\n", model.Version)
			fmt.Printf("  DisplayName: %s\n", model.DisplayName)
			fmt.Printf("  Description: %s\n", model.Description)
			fmt.Printf("  InputTokenLimit: %d\n", model.InputTokenLimit)
			fmt.Printf("  OutputTokenLimit: %d\n", model.OutputTokenLimit)
			fmt.Printf("  SupportedActions: %v\n", model.SupportedActions)
			fmt.Println()
		}
		return
	}

	// Add "models/" prefix if not already present
	modelName := *model
	if !strings.HasPrefix(modelName, "models/") {
		modelName = "models/" + modelName
	}
	log.Printf("Model: %s", modelName)

	// Create a model instance
	if client.ClientConfig().Backend == genai.BackendVertexAI {
		fmt.Println("Calling VertexAI Backend...")
	} else {
		fmt.Println("Calling GeminiAPI Backend...")
	}

	var config *genai.GenerateContentConfig = &genai.GenerateContentConfig{Temperature: genai.Ptr[float32](0.5)}

	// Create a new Chat.
	chat, err := client.Chats.Create(ctx, modelName, config, nil)
	if err != nil {
		log.Fatalf("Failed to create chat: %v", err)
	}

	// Try streaming API since native audio models support bidiGenerateContent
	fmt.Println("Sending first message (streaming)...")
	for result, err := range chat.SendMessageStream(ctx, genai.Part{Text: "What's the weather in San Francisco?"}) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Text())
	}
	fmt.Println()

	// Send second chat message.
	fmt.Println("Sending second message (streaming)...")
	for result, err := range chat.SendMessageStream(ctx, genai.Part{Text: "How about New York?"}) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Text())
	}
	fmt.Println()
}

func main() {
	// Get API key from environment variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is not set")
	}

	ctx := context.Background()
	flag.Parse()
	chat(ctx, apiKey)
}
