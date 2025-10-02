package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type SqlIndex struct {
	db             *sql.DB
	stmtInsertWord *sql.Stmt
	stmtInsertUrl  *sql.Stmt
}

func NewSqlIndex() *SqlIndex {
	db := sql.OpenDB("sqlite3")

	sqlInsertWord := `INSERT INTO words(word) IF NOT EXISTS VALUES(?);`
	stmtInsertWord, err := db.Prepare(sqlInsertWord)
	if err != nil {
		// handle
	}

	sqlInsertUrl := `INSERT INTO urls(url) VALUES(?);`
	stmtInsertUrl, err := db.Prepare(sqlInsertUrl)
	if err != nil {
		// handle
	}
	return &SqlIndex{
		db:             db,
		stmtInsertWord: stmtInsertWord,
		stmtInsertUrl:  stmtInsertUrl,
	}
}

func (s *SqlIndex) Insert(words []string, url string) {

	// Transactions give us a way to prevent partial commits
	// to the database, for ACID consistency
	// xactions also give us a performance benefit by committing
	// all the INSERTs in a batch
	tx, err := s.db.Begin()
	if err != nil {
		// handle error
	}
	defer tx.Commit()

	// For ORM users, transaction are automatic, but only if you
	// insert a slice of objects.
	// s.db.Create(&words)

	tx.Stmt(s.stmtInsertUrl).Exec(url)

	// Use a hash table to get the property where each
	// key can exist only once, so we end up with keys = words in doc,
	// values = count of each of those words.
	// THis relies on hash table calc being much faster than disk hits
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for word, count := range wordMap {
		tx.Stmt(s.stmtInsertWord).Exec(word)
		// update the frequency/hits table using count
	}

}
