package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	sqlite_vec "github.com/asg017/sqlite-vec-go-bindings/cgo"
	_ "github.com/mattn/go-sqlite3"
	openai "github.com/sashabaranov/go-openai"
)

func CreateBlob(client *openai.Client, plain string) []byte {
	req := openai.EmbeddingRequest{
		Model: openai.LargeEmbedding3,
		Input: plain,
	}

	resp, err := client.CreateEmbeddings(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	blob, err := sqlite_vec.SerializeFloat32(resp.Data[0].Embedding)
	if err != nil {
		log.Fatal(err)
	}
	return blob
}

func main() {
	sqlite_vec.Auto()
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`
		CREATE VIRTUAL TABLE test USING vec0(
			id INTEGER PRIMARY KEY,
			plain TEXT,
			embed FLOAT[3072]);
	`)
	if err != nil {
		log.Fatal(err)
	}
	client := openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY"))

	people := []string{
		"Greg Benson",
		"Mehmet Emre",
		"Ellen Veomett",
		"Kelsey Urgo",
		"Phil Peterson",
	}
	for idx, person := range people {
		blob := CreateBlob(client, person)

		_, err = db.Exec(`
		INSERT INTO test VALUES(?, ?, ?);
	`, idx+1, person, blob)
		if err != nil {
			log.Fatal(err)
		}
	}
	blob := CreateBlob(client, "elln viomett")

	rows, err := db.Query(`
		SELECT id, plain, distance 
		FROM test WHERE embed MATCH ? ORDER BY distance LIMIT 3;
	`, blob)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int64
		var plain string
		var distance float32
		err = rows.Scan(&id, &plain, &distance)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("id: %d plain: %s distance: %f\n", id, plain, distance)
	}
}
