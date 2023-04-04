package aws

import (
	"context"
	"fmt"
	"github.com/huseyinbabal/aiws/internal/chatgpt"
	"strings"
)

type Client struct {
	ChatGpt *chatgpt.Client
}

func New(chatGpt *chatgpt.Client) *Client {
	return &Client{ChatGpt: chatGpt}
}

func (c *Client) GenerateCommand(ctx context.Context, statement string) (string, error) {
	q := fmt.Sprintf("Could you please show me code example to %s in json format with AWS CLI? Only respond with code as plain text without code block syntax around it.", statement)
	res, err := c.ChatGpt.Ask(ctx, q)
	if err != nil {
		return "", err
	}
	return trimUnwantedChars(res), nil
}

func trimUnwantedChars(command string) string {
	command = strings.ReplaceAll(command, "`", "")
	return command
}
