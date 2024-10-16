package main

import (
	"fmt"
)

func cariFaktor(b int) []int {
	var faktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func cekPrima(b int) bool {
	// Bilangan prima hanya punya 2 faktor (1 dan dirinya sendiri)
	faktor := cariFaktor(b)
	return len(faktor) == 2
}

func main() {
	var b int

	// Input nilai b
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&b)

	// Mencari faktor dari b
	faktor := cariFaktor(b)

	// Menampilkan faktor-faktor
	fmt.Printf("Bilangan: %d\n", b)
	fmt.Print("Faktor: ")
	for _, f := range faktor {
		fmt.Print(f, " ")
	}
	fmt.Println()

	// Menentukan apakah b adalah bilangan prima
	if cekPrima(b) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
