package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func getGeneratedText(prompt string) {
	//Get gemini api client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatalf("Failed to get client: %v", err)
	}
	defer client.Close()

	//Send prompt to gemini and get generated response
	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("Failed to get response from gemini: %v", err)
	}

	//Output response
	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			fmt.Println(part)
		}
	}
}

func main() {
	vocabularyList := [3]string{"elucidate", "languishing", "bluntly"}
	prompt := fmt.Sprintf("Create a English sentence, using following words: %s, %s, %s",
		vocabularyList[0], vocabularyList[1], vocabularyList[2])

	fmt.Println("")
	fmt.Println("")

	fmt.Println("++++++ Prompt ++++++")
	fmt.Println(prompt)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("++++++ Generated response ++++++")
	getGeneratedText(prompt)

	fmt.Println("")
	fmt.Println("")
}
