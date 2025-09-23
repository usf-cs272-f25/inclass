package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	name := flag.String("name", "database.db", "name of the database")
	flag.Parse()
	db, err := sql.Open("sqlite3", *name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cars (
			id INTEGER PRIMARY KEY,
			make TEXT,
			model TEXT
		);
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`
		INSERT INTO cars(make, model) VALUES("Renault", "2CV")
	`)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query(`
		SELECT (model) FROM cars where id = 1;
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var model string
		if err = rows.Scan(&model); err == nil {
			fmt.Println("model: ", model)
		}
	}

}
