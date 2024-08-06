package main

import (
	"context"
	"fmt"
	"log"
	"magic-rules-qa/parser"
	"net/url"
	"os"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/embeddings/jina"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

func create_vectorstore() (qdrant.Store, error) {
	if jinakey := os.Getenv("JINA_API_KEY"); jinakey == "" {
		fmt.Errorf("JINA_API_KEY not set")
	}

	j, err := jina.NewJina(jina.WithModel("jina-embeddings-v2-base-en"))
	if err != nil {
		log.Fatal(err)
	}

	e, err := embeddings.NewEmbedder(j)
	if err != nil {
		log.Fatal(err)
	}

	url, err := url.Parse("http://localhost:6333")
	if err != nil {
		log.Fatal(err)
	}

	store, err := qdrant.New(
		qdrant.WithURL(*url),
		qdrant.WithCollectionName("magic-qa"),
		qdrant.WithEmbedder(e),
	)

	return store, nil
}

func load_docs() ([]schema.Document, error) {
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

func load_documents_to_db(ctx context.Context, store qdrant.Store) {
	inserted_docs, err := load_docs()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("loading documents on the database")
	_, err = store.AddDocuments(ctx, inserted_docs)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	store, err := create_vectorstore()
	if err != nil {
		log.Fatal(err)
	}

	load_documents_to_db(ctx, store)

	docs, err := store.SimilaritySearch(ctx, "What happens when I saddle a creature?", 5)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range docs {
		log.Println(d.Metadata)
		log.Println(d.PageContent)
		log.Println(d.Score)
	}
}
