package vectorstore

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/embeddings/jina"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

func NewQdrant() (qdrant.Store, error) {
	if jinakey := os.Getenv("JINA_API_KEY"); jinakey == "" {
		log.Fatal("JINA_API_KEY not set")
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

	hasCollection, err := getCollection(*url, "magic-qa")
	if err != nil {
		log.Fatal(err)
	}

	if !hasCollection {
		err = newCollection(*url, "magic-qa")
		if err != nil {
			log.Fatal(err)
		}

	}

	store, err := qdrant.New(
		qdrant.WithURL(*url),
		qdrant.WithCollectionName("magic-qa"),
		qdrant.WithEmbedder(e),
	)
	if err != nil {
		log.Fatal(err)
	}

	return store, nil
}

func getCollection(url url.URL, name string) (bool, error) {
	url.Path = fmt.Sprintf("/collections/%s", name)
	jsonData := `{"vectors": {"size": 768, "distance": "Cosine"}}`

	reqGet, err := http.NewRequest(http.MethodGet, url.String(), bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	respGet, err := client.Do(reqGet)
	if err != nil {
		return false, err
	}

	if respGet.StatusCode != 200 {
		return false, nil
	}

	return true, nil
}

func newCollection(url url.URL, name string) error {
	url.Path = fmt.Sprintf("/collections/%s", name)
	jsonData := `{"vectors": {"size": 768, "distance": "Cosine"}}`

	reqPut, err := http.NewRequest(http.MethodPut, url.String(), bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return err
	}

	reqPut.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqPut)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	return nil
}
