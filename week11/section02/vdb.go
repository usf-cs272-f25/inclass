package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	sqlite_vec "github.com/asg017/sqlite-vec-go-bindings/cgo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sashabaranov/go-openai"
)

type VectorDB struct {
	client *openai.Client
	db     *sql.DB
}

func NewVectorDB() *VectorDB {
	sqlite_vec.Auto()
	name := "demo.db"
	os.Remove(name)
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}

	return &VectorDB{
		client: openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY")),
		db:     db,
	}
}

func (v *VectorDB) CreateTable() {
	_, err := v.db.Exec(`
		CREATE VIRTUAL TABLE demo USING vec0(
			id INTEGER PRIMARY KEY,
			embedding FLOAT[3072]
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func (v *VectorDB) CreateBlob(s string) []byte {
	req := openai.EmbeddingRequest{
		Input: s,
		Model: openai.LargeEmbedding3,
	}

	resp, err := v.client.CreateEmbeddings(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	blob, err := sqlite_vec.SerializeFloat32(resp.Data[0].Embedding)
	if err != nil {
		log.Fatal(err)
	}
	return blob
}

func (v *VectorDB) Insert(id int, blob []byte) {
	_, err := v.db.Exec(`
		INSERT INTO demo VALUES (?, ?);
	`, id, blob)
	if err != nil {
		log.Fatal(err)
	}
}

func (v *VectorDB) Query(q string) {
	// Remember to create the embedding blob for the
	// text you want to query for. This is how sqlite-vec
	// calculates the nearest neighbors
	blob := v.CreateBlob(q)

	rows, err := v.db.Query(`
		SELECT id, distance FROM demo WHERE embedding
		MATCH ? ORDER BY distance LIMIT 3;
	`, blob)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var distance float32
		rows.Scan(&id, &distance)
		log.Printf("found ID: %d, distance: %f", id, distance)
	}
}
