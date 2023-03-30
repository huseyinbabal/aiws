package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/alecthomas/chroma/quick"
	"github.com/fatih/color"
	"github.com/huseyinbabal/aiws/internal/aws"
	"github.com/huseyinbabal/aiws/internal/chatgpt"
	"golang.design/x/clipboard"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	apiKey := os.Getenv("OPEN_AI_API_KEY")
	chatGptClient := chatgpt.New(apiKey)
	awsClient := aws.New(chatGptClient)

	statement := os.Args[1:]
	command, err := awsClient.GenerateCommand(context.Background(), strings.Join(statement, " "))
	if err != nil {
		log.Fatalf("failed to generate command. %v", err)
	}

	err = clipboard.Init()
	if err != nil {
		log.Fatalf("failed to init clipboard. %v", err)
	}
	clipboard.Write(clipboard.FmtText, []byte(command))
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice("👇 Command is copied to your clipboard")

	quick.Highlight(os.Stdout, command, "json", "terminal256", "monokai")
	fmt.Println()
	notice("Would you like to run ☝️ command? (y/N)")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	if input.Text() == "y" {
		doExec(command)
	}
}

func doExec(command string) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()

	if err != nil {
		quick.Highlight(os.Stdout, errb.String(), "bash", "text", "terminal256")
	} else {
		err := quick.Highlight(os.Stdout, outb.String(), "json", "terminal256", "monokai")
		if err != nil {
			log.Fatal(err)
		}
	}
}
