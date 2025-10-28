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

	db, err := sql.Open("sqlite3", "demo.db")
	if err != nil {
		log.Fatal(err)
	}

	return &VectorDB{
		client: openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY")),
		db:     db,
	}
}

func (v *VectorDB) CreateTable() {
	// TODO: double check FLOAT vs byte?
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

func (v *VectorDB) CreateEmbedding(s string) []byte {
	req := openai.EmbeddingRequest{
		Input: s,
		Model: openai.LargeEmbedding3,
	}

	resp, err := v.client.CreateEmbeddings(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}

	bts, err := sqlite_vec.SerializeFloat32(resp.Data[0].Embedding)
	if err != nil {
		log.Fatal(err)
	}
	return bts
}

func (v *VectorDB) Insert(id int, blob []byte) {
	_, err := v.db.Exec(`
		INSERT INTO demo VALUES (?, ?);
	`, id, blob)
	if err != nil {
		log.Fatal(err)
	}
}
