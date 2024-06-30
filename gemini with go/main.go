package main

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyDoK9M4MP1hRvg51SNO0p1_uRT2zLJfhkHkYXV"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with both text-only and multimodal prompts
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("essay on global warming"))

	if resp != nil {
		candidates := resp.Candidates
		if candidates != nil {
			for _, candidate := range candidates {
				content := candidate.Content
				if content != nil {
					text := content.Parts[0]
					log.Println("Candidate Content Text: ", text)
				}
			}
		} else {
			log.Println("Candidate slice is nil")
		}
	} else {
		log.Println("Response is nil")
	}
	if err != nil {
		log.Fatal(err)
	}
}
