package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

var model = flag.String("model", "gemini-2.0-flash", "the model name, e.g. gemini-2.0-flash")

func chat(ctx context.Context, apiKey string) {
	// Create a client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	log.Printf("Model: %s", *model)

	// Create a model instance
	if client.ClientConfig().Backend == genai.BackendVertexAI {
		fmt.Println("Calling VertexAI Backend...")
	} else {
		fmt.Println("Calling GeminiAPI Backend...")
	}

	var config *genai.GenerateContentConfig = &genai.GenerateContentConfig{Temperature: genai.Ptr[float32](0.5)}

	// Create a new Chat.
	chat, err := client.Chats.Create(ctx, *model, config, nil)
	if err != nil {
		log.Fatalf("Failed to create chat: %v", err)
	}

	// Send first chat message.
	result, err := chat.SendMessage(ctx, genai.Part{Text: "What's the weather in San Francisco?"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())

	// Send second chat message.
	result, err = chat.SendMessage(ctx, genai.Part{Text: "How about New York?"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
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
