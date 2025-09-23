package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS fruits (
			id INTEGER PRIMARY KEY,
			name TEXT,
			color TEXT
		);
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`
		INSERT INTO fruits(name, color) VALUES("banana", "yellow");
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query(`
		SELECT name FROM fruits WHERE id = 1;
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("name: ", name)
	}
}
