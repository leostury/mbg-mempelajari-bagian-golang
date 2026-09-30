package main

import "fmt"

func tutorial() {
	// variable
	var name string = "leo"
	age := 21
	fmt.Printf("hello %s umur %d\n", name, age)

	// array
	hobbies := []string{"jalan", "berenang", "game"}
	fmt.Printf("hobbie saya")
	for _, hobby := range hobbies {
		fmt.Printf(" %s\n", hobby)
	}

	favoriteFood := make([]string, 0)

	favoriteFood = append(favoriteFood, "Soto")
	favoriteFood = append(favoriteFood, "Nasi Goreng")

	fmt.Println("makanan favorite saya")
	for _, food := range favoriteFood {
		fmt.Printf("\t- %s\n", food)
	}

	// array object
	animals := make([]Animal, 0)
	animals = append(animals, Animal{
		Type:  "kucing",
		Name:  "mimi",
		Color: "oranye",
	})
	animals = append(animals, Animal{
		Type:  "anjing",
		Name:  "doggy",
		Color: "hitam",
	})
	animals = append(animals, Animal{
		Type:  "burung",
		Name:  "pipit",
		Color: "biru",
	})

	fmt.Println("hewan :")
	for _, animal := range animals {
		fmt.Printf("jenis %s, nama %s, warna %s\n", animal.Type, animal.Name, animal.Color)
	}

}

type Animal struct {
	Type  string
	Name  string
	Color string
}

// GET: ambil data
// POST: insert atau submit
// PUT: update data
// PATCH: update data, mirip PUT, tapi cuman 1 key doank (is_active)
// DELETE: hapus data

// user (mobile app) <-> api (backend) <-> router (endpoint) <-> handler (logic) <-> model (table database)
