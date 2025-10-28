package main

func main() {
	db := NewVectorDB()
	db.CreateTable()
	blob := db.CreateEmbedding("who teaches PHIL 230?")
	db.Insert(1, blob)
}
