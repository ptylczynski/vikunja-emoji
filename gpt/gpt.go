package gpt

import (
	"context"
	"fmt"
	"github.com/forPelevin/gomoji"
	"github.com/sashabaranov/go-openai"
	"os"
)

func GetEmoji(title string) string {
	emoji := fetchEmoji(title)
	if gomoji.RemoveEmojis(emoji) == "" {
		fmt.Println("Generated emoji for title " + title + " is " + emoji)
		return emoji
	} else {
		fmt.Println("Response contained not only emojis")
		return ""
	}
}

func fetchEmoji(title string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Podaj emoji dla tekstu: " + title,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}
	content := resp.Choices[0].Message.Content
	fmt.Println("Completion: " + content)
	return content
}
