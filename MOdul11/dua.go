package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Berapa Banyak Ikan yang Dijual: ")
	fmt.Scan(&x)
	fmt.Print("Berapa Banyak Ikan Untuk Ditempatkan di Wadah: ")
	fmt.Scan(&y)

	if x <= 0 || x > 1000 {
		fmt.Println("Jumlah Ikan Tidak Boleh Lebih Dari 1000")
		return
	}
	if y <= 0 || y > 1000 {
		fmt.Println("Ikan Melebihi Kapasitas Wadah!")
		return
	}


	fmt.Println("Masukkan Berat Ikan:")
	weights := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	totalContainers := (x + y - 1) / y 
	containerWeights := make([]float64, totalContainers)
	for i := 0; i < x; i++ {
		containerWeights[i/y] += weights[i]
	}

	fmt.Println("Total Berat di Setiap Wadah:")
	for _, weight := range containerWeights {
		fmt.Printf("%.2f ", weight)
	}
	fmt.Println()


	var totalWeight float64
	for _, weight := range containerWeights {
		totalWeight += weight
	}
	averageWeight := totalWeight / float64(totalContainers)
	fmt.Printf("Rata-rata Berat di Setiap Wadah: %.2f\n", averageWeight)
}
