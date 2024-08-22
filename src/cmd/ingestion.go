package cmd

import (
	"context"
	"fmt"
	"log"
	"magic-rules-qa/parser"
	"magic-rules-qa/vectorstore"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

func getParsedDocuments() ([]schema.Document, error) {
	var docs []schema.Document

	rule, _, err := parser.ParseFile("../data/MagicCompRules_20240607.txt")
	if err != nil {
		fmt.Println("error parsing files", err)
	}

	for _, v := range rule {
		docs = append(docs, schema.Document{PageContent: v.Text, Metadata: map[string]any{"code": v.Code}})
	}

	log.Println("Documents split: ", len(docs))

	return docs, nil
}

func loadDocuments(ctx context.Context, store qdrant.Store, documents []schema.Document) {
	log.Println("loading documents on the database")
	_, err := store.AddDocuments(ctx, documents)
	if err != nil {
		log.Fatal(err)
	}
}

func Ingestion() *cobra.Command {
	return &cobra.Command{
		Use:   "ingestion of the data",
		Short: "Make ingestion of the data",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			store, err := vectorstore.NewQdrant()
			if err != nil {
				log.Fatal(err)
			}

			docs, err := getParsedDocuments()
			if err != nil {
				log.Fatal(err)
			}

			loadDocuments(ctx, store, docs)
		},
	}
}
