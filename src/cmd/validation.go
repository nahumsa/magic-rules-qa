package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"magic-rules-qa/vectorstore"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

type validationExample struct {
	Question string `json:"question"`
	Rule     string `json:"rule"`
}

func getValidationExamples() ([]validationExample, error) {
	jsonFile, err := os.Open("../data/query_validation.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var examples []validationExample
	json.Unmarshal(byteValue, &examples)
	return examples, nil
}

func makeSearch(ctx context.Context, store qdrant.Store, query string, numDocs int) ([]schema.Document, error) {
	docs, err := store.SimilaritySearch(ctx, query, numDocs)
	if err != nil {
		log.Fatal(err)
	}

	return docs, nil
}

func makeValidation(documents []schema.Document, example validationExample) (float64, error) {
	hits := 0

	for _, doc := range documents {
		code, ok := doc.Metadata["code"].(string)

		if ok {
			if strings.HasPrefix(code, example.Rule) {
				hits++
			}
		}
	}

	return float64(hits) / float64(len(documents)), nil
}

func Validation() *cobra.Command {
	return &cobra.Command{
		Use:   "validation for retrieval",
		Short: "Validation for retrieval from the database",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			store, err := vectorstore.NewQdrant()
			if err != nil {
				log.Fatal(err)
			}

			examples, err := getValidationExamples()
			if err != nil {
				log.Fatal(err)
			}

			for _, ex := range examples {
				docs, err := makeSearch(ctx, store, ex.Question, 5)
				if err != nil {
					log.Fatal(err)
				}

				hits, err := makeValidation(docs, ex)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("question: %s hits: %.2f%%\n", ex.Question, hits)

			}
		},
	}
}
