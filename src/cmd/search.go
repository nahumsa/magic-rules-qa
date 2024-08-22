package cmd

import (
	"context"
	"log"
	"magic-rules-qa/vectorstore"

	"github.com/spf13/cobra"
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
		},
	}
}
