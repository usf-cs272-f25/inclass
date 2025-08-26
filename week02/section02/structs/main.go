package main

import "fmt"

type Flower struct {
	Name      string
	NumPetals int
}

func findFlower(b []Flower, name string) *Flower {
	for idx := range b {
		if b[idx].Name == name {
			retval := Flower{
				b[idx].Name,
				b[idx].NumPetals,
			}

			// &retval is on the stack, but Go compiler
			// treats it as a "stack escape" and creates
			// a heap object for the Flower
			return &retval
		}
	}
	return nil
}

func printFlower(f *Flower) {
	// When using a pointer, '.' dereferences the ptr
	fmt.Printf("Name: %s, NumPetals: %d\n", f.Name, f.NumPetals)

	//f.Name = "Rhododendron"
}

func main() {
	bouquet := []Flower{
		{"rose", 5},
		{"lily", 3},
		{"clover", 12},
	}

	printFlower(&bouquet[0])
	f := findFlower(bouquet, "rose")
	if f == nil {
		fmt.Println("didn't find any roses")
	} else {
		fmt.Println(f)
	}

}
