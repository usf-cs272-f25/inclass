package main

func main() {
	db := NewVectorDB()
	db.CreateTable()

	people := []string{
		"Phil Peterson",
		"Ellen Veomett",
		"Greg Benson",
		"Mehmet Emre",
		"Matthew Malensek",
		"Kelsey Urgo",
	}

	for idx, p := range people {
		blob := db.CreateBlob(p)
		db.Insert(idx+1, blob)
	}

	p := "Gregory Benson"
	db.Query(p)
}
