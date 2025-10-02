package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type SqlIndex struct {
	db *sql.DB
}

func (s *SqlIndex) Insert(words []string, url string) {
	sqlInsert := `INSERT INTO words (word, count) VALUES(?)`
	sqlQuery := `SELECT id FROM words WHERE word = ?`

	// This is pretty slow, hitting the DB twice for each word
	// How can we make fewer DB hits?

	// Database transactions can help "batch" operations, as well
	// as provide a "rollback" path if operations fail

	tx, err := s.db.Begin()
	if err != nil {
		// handle the error
	}
	defer tx.Commit()

	// Use a map, making use of the property of maps that
	// each key exists only once, so we turn 30 Romeo inserts
	// into one Romeo insert

	m := make(map[string]int)
	for idx, word := range words {
		if val, exists := m[word]; exists {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	for word, _ := range m {
		err := s.db.Exec(sqlInsert, word)
		if err != nil {
			// Catch errors like schema constraint violations
			// and roll back the transaction rather than allowing
			// partial completion of the transaction - see ACID
			tx.Rollback()
			return
		}
		rows := s.db.Query(sqlQuery, word)
		for rows.Next() {
			var id int
			rows.Scan(&id)
		}
	}
}
