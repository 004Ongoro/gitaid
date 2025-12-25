package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/004Ongoro/gitaid/internal/config"
	"github.com/004Ongoro/gitaid/internal/git"
	"google.golang.org/genai"
)

const systemPrompt = `You are a git commit generator. 
Analyze the provided diff and write a detailed commit message following the Conventional Commits specification.
Format:
<type>(<scope>): <subject>

[optional body]

[optional footer(s)]

Types: feat, fix, docs, style, refactor, perf, test, chore.
Do not use emojis. Keep the subject line under 50 characters. 
The body should explain the 'why' of the change, not just the 'what'.`

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	diff, err := git.GetStagedDiff()
	if err != nil {
		log.Fatalf("Git error: %v", err)
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: cfg.GeminiKey,
	})
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	generateMessage(ctx, client, cfg, diff)
}

func generateMessage(ctx context.Context, client *genai.Client, cfg *config.Config, diff string) {
	fmt.Println("Analyzing changes and generating commit message...")

	prompt := fmt.Sprintf("%s\n\nDiff to analyze:\n%s", systemPrompt, diff)

	result, err := client.Models.GenerateContent(ctx, cfg.Model, genai.Text(prompt), nil)
	if err != nil {
		log.Fatalf("AI generation failed: %v", err)
	}

	//  extract text from the Candidate's Parts
	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		var parts []string
		for _, part := range result.Candidates[0].Content.Parts {
			// Extract just the text value from the part
			parts = append(parts, fmt.Sprintf("%s", part.Text))
		}
		commitMsg := strings.Join(parts, "\n")

		fmt.Println("\n--- Suggested Commit Message ---")
		fmt.Println(commitMsg)
		fmt.Println("--------------------------------")

		handleUserChoice(ctx, client, cfg, diff, commitMsg)
	} else {
		fmt.Println("AI could not generate a message for these changes.")
	}
}

func handleUserChoice(ctx context.Context, client *genai.Client, cfg *config.Config, diff, currentMsg string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nOptions: (a)ccept & commit, (r)egenerate, (q)uit: ")
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		switch input {
		case "a":
			fmt.Println("\nCommitting changes...")
			// Execute git commit -m "message"
			cmd := exec.Command("git", "commit", "-m", currentMsg)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Fatalf("Failed to execute git commit: %v", err)
			}

			fmt.Println("Commit successful!")
			os.Exit(0)
		case "r":
			fmt.Println("")
			generateMessage(ctx, client, cfg, diff)
			return
		case "q":
			fmt.Println("Discarding message.")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose a, r, or q.")
		}
	}
}
