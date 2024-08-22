package cmd

import (
	"context"
	"fmt"
	"log"
	"magic-rules-qa/vectorstore"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func Search() *cobra.Command {
	return &cobra.Command{
		Use:   "search for a phrase",
		Short: "Search for a given phrase",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			store, err := vectorstore.NewQdrant()
			if err != nil {
				log.Fatal(err)
			}

			docs, err := store.SimilaritySearch(ctx, args[0], 5)
			if err != nil {
				log.Fatal(err)
			}

			for _, d := range docs {
				log.Println(d.Metadata)
				log.Println(d.PageContent)
				log.Println(d.Score)
			}

			apiKey := os.Getenv("GOOGLE_API_KEY")
			llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey), googleai.WithDefaultModel("gemini-1.5-flash"))
			if err != nil {
				log.Fatal(err)
			}

			content := []llms.MessageContent{
				llms.TextParts(llms.ChatMessageTypeSystem, "You are an expert in Magic: The Gathering"),
				llms.TextParts(llms.ChatMessageTypeHuman, "Given the following rule: %s Answer the question: %s", docs[0].PageContent, args[0]),
			}

			completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				fmt.Print(string(chunk))
				return nil
			}))
			if err != nil {
				log.Fatal(err)
			}
			_ = completion
		},
	}
}
