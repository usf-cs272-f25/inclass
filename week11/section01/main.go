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

type DB struct {
	db     *sql.DB
	client *openai.Client
}

func NewDB() *DB {
	// Load the vec0 vector extension
	sqlite_vec.Auto()
	db, err := sql.Open("sqlite3", "demo.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create a SQLite virtual table using the vector extension
	_, err = db.Exec(`
		CREATE VIRTUAL TABLE demo USING vec0(
			id INTEGER PRIMARY KEY,
			embedding FLOAT[3072]
		);
	`)
	log.Fatal(err)

	return &DB{
		db:     db,
		client: openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY")),
	}
}

func (db *DB) CreateEmbedding(input string) []float32 {
	req := openai.EmbeddingRequest{
		Input: "Who is teaching CS 315",
		Model: openai.LargeEmbedding3,
	}
	resp, err := db.client.CreateEmbeddings(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Data[0].Embedding
}

func (db *DB) InsertEmbedding(rowid int, e []float32) {
	// Insert the embedding we got from OpenAI into the virtual table
	_, err := db.db.Exec(`
		INSERT INTO demo VALUES(?, ?);
	`, rowid, e)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) Query(p string) {
	// e := db.CreateEmbedding(p)
	// sql
	// rows, err := db.db.Query()
}

func main() {
	courses := []string{
		"Phil Peterson teaches CS 272",
		"Greg Benson teaches CS 315",
	}
	db := NewDB()
	for idx, c := range courses {
		e := db.CreateEmbedding(c)
		db.InsertEmbedding(idx, e)
	}

	// db.Query("who teaches CS 315?")
}
