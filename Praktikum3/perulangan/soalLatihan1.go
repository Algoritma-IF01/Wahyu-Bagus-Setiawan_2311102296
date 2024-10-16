package main

import (
	"fmt"
)

func main() {
	// Definisikan warna yang benar
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	// Lakukan 5 percobaan
	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d\n", i)
		fmt.Printf("Masukkan warna pertama : ")
		fmt.Scanln(&warna1)
		fmt.Printf("Masukkan warna kedua : ")
		fmt.Scanln(&warna2)
		fmt.Printf("Masukkan warna ketiga : ")
		fmt.Scanln(&warna3)
		fmt.Printf("Masukkan warna keempat : ")
		fmt.Scanln(&warna4)

		// Periksa apakah urutan warna sesuai
		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] || warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {
			hasil = false
		}
	}

	// Tampilkan hasil
	fmt.Println("BERHASIL:", hasil)
}
