package main

import (
	"fmt"
	"math"
)

func main() {
	var totalKantong1, totalKantong2 float64
	var beratKantong1, beratKantong2 float64

	for {
		fmt.Print("Masukkan berat belanjaan di kantong Pertama: ")
		fmt.Scan(&beratKantong1)

		// Meminta input untuk kantong kedua
		fmt.Print("Masukkan berat belanjaan di kantong Kedua: ")
		fmt.Scan(&beratKantong2)

		// Cek jika berat negatif
		if beratKantong1 < 0 || beratKantong2 < 0 {
			fmt.Println("Berat tidak boleh negatif. Program selesai.")
			return
		}

		totalKantong1 += beratKantong1
		totalKantong2 += beratKantong2

		// Cek jika total berat lebih dari 150 kg
		if totalKantong1+totalKantong2 > 150 {
			fmt.Println("Total berat melebihi 150 kg. Program selesai.")
			return
		}

		// Cek selisih berat kedua kantong
		if math.Abs(totalKantong1-totalKantong2) >= 9 {
			fmt.Println("Sepeda Motor Pak Andi Akan Oleng: True")
		} else {
			fmt.Println("Sepeda Motor Pak Andi Akan Oleng: False")
		}
	}
}
