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
	name := "demo.db"
	os.Remove(name)
	sqlite_vec.Auto()
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}

	// Create a SQLite virtual table using the vector extension
	_, err = db.Exec(`
		CREATE VIRTUAL TABLE demo USING vec0(
			id INTEGER PRIMARY KEY,
			plain TEXT,
			embedding FLOAT[3072]
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		db:     db,
		client: openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY")),
	}
}

func (db *DB) CreateBlob(input string) []byte {
	req := openai.EmbeddingRequest{
		Input: "Who is teaching CS 315",
		Model: openai.LargeEmbedding3,
	}
	resp, err := db.client.CreateEmbeddings(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	bts, err := sqlite_vec.SerializeFloat32(resp.Data[0].Embedding)
	if err != nil {
		log.Fatal(err)
	}
	return bts
}

func (db *DB) InsertBlob(rowid int, p string, b []byte) {
	// Insert the embedding we got from OpenAI into the virtual table
	_, err := db.db.Exec(`
		INSERT INTO demo VALUES(?, ?, ?);
	`, rowid, p, b)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) Query(p string) {
	b := db.CreateBlob(p)
	rows, err := db.db.Query(`
		SELECT id, plain, distance FROM demo WHERE embedding MATCH ? ORDER BY distance LIMIT 3;
	`, b)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int32
		var distance float32
		var plain string
		err = rows.Scan(&id, &plain, &distance)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("id: %d, plain: %s, distance: %f\n", id, plain, distance)
	}
}

func main() {
	people := []string{
		"Phil Peterson",
		"Greg Benson",
		"Ellen Veomett",
		"EJ Jung",
		"Mehmet Emre",
	}
	db := NewDB()
	for idx, p := range people {
		b := db.CreateBlob(p)
		db.InsertBlob(idx, p, b)
	}

	db.Query("Philip Peterson")
}
