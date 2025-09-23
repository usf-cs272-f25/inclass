package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// const (
// 	white = iota
// 	orange
// )

func main() {

	name := flag.String("name", "", "name of the database")
	flag.Parse()

	db, err := sql.Open("sqlite3", *name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS chippers(
			id INTEGER PRIMARY KEY,
			brand TEXT,
			year INTEGER,
			color INTEGER);
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(`
		INSERT INTO chippers(brand, year, color) VALUES("Deere", 2022, 1);
	`)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query(`
		SELECT brand FROM chippers WHERE id = 1;
	`)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var brand string
		err = rows.Scan(&brand)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
