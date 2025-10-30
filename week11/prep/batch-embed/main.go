package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	sqlite_vec "github.com/asg017/sqlite-vec-go-bindings/cgo"
	_ "github.com/mattn/go-sqlite3"
	openai "github.com/sashabaranov/go-openai"
)

func CreateBlobs(client *openai.Client, plain []string) [][]byte {
	req := openai.EmbeddingRequest{
		Model: openai.LargeEmbedding3,
		Input: plain,
	}

	resp, err := client.CreateEmbeddings(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	blobs := make([][]byte, len(plain))
	for i := range len(plain) {
		blobs[i], err = sqlite_vec.SerializeFloat32(resp.Data[i].Embedding)
		if err != nil {
			log.Fatal(err)
		}
	}
	return blobs
}

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY"))

	people := []string{
		"Greg Benson",
		"Mehmet Emre",
		"Ellen Veomett",
		"Kelsey Urgo",
		"Phil Peterson",
		"David Guy Brizan",
		"EJ Jung",
		"David Wolber",
		"Matthew Malensek",
		"Paul Haskell",
	}

	start := time.Now()
	// for _, p := range people {
	// 	CreateBlobs(client, []string{p})
	// }
	CreateBlobs(client, people)

	fmt.Println(time.Since(start))

}
