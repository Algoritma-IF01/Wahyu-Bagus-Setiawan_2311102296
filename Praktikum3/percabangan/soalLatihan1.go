package main

import (
	"fmt"
)

func hitungBiayaPos(berat int) (int, int, int) {
	kg := berat / 1000             
	gram := berat % 1000           
	biayaDasar := kg * 10000       
	biayaTambahan := 0             

	if gram > 0 {
		if gram < 500 {
			biayaTambahan = gram * 5 
		} else {
			biayaTambahan = gram * 15 
		}
	}

	totalBeratDigit := kg + gram/100 + (gram%100)/10 + gram%10
	if kg == 0 && totalBeratDigit >= 10 {
		biayaDasar = 10000 
		biayaTambahan = 0  
	}

	totalBiaya := biayaDasar + biayaTambahan

	return kg, gram, totalBiaya
}

func main() {
	var berat int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scanln(&berat)

	kg, gram, totalBiaya := hitungBiayaPos(berat)

	fmt.Printf("Berat parsel (gram): %d\n", berat)
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
