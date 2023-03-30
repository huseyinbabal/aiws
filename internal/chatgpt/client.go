package chatgpt

import (
	"context"
	"github.com/PullRequestInc/go-gpt3"
	"strings"
)

type Client struct {
	gpt gpt3.Client
}

func New(apiKey string) *Client {
	c := gpt3.NewClient(apiKey)
	return &Client{
		gpt: c,
	}
}

func (c *Client) Ask(ctx context.Context, question string) (string, error) {
	sb := strings.Builder{}
	err := c.gpt.CompletionStreamWithEngine(ctx,
		gpt3.TextDavinci003Engine,
		gpt3.CompletionRequest{
			Prompt:      []string{question},
			MaxTokens:   gpt3.IntPtr(300),
			Temperature: gpt3.Float32Ptr(0),
		},
		func(response *gpt3.CompletionResponse) {
			text := response.Choices[0].Text
			sb.WriteString(text)
		},
	)
	if err != nil {
		return "", err
	}
	response := sb.String()
	response = strings.TrimLeft(response, "\n")
	return response, nil
}
