package main

import (
	"fmt"
	"math"
)

func hitung_f(k int) float64 {
	f_k := math.Pow(float64(4*k+2), 2) / (float64(4*k+1) * float64(4*k+3))
	return f_k
}

func hitung_akar2(K int) float64 {
	akar_2 := 1.0
	for k := 0; k < K; k++ {
		akar_2 *= math.Pow(float64(4*k+2), 2) / (float64(4*k+1) * float64(4*k+3))
	}
	return akar_2
}

func main() {
	var K int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scanln(&K)

	nilai_f := hitung_f(K)
	fmt.Printf("Nilai K = %d\n", K)
	fmt.Printf("Nilai f(K) = %.10f\n", nilai_f)

	nilai_akar2 := hitung_akar2(K)
	fmt.Printf("Nilai akar 2 = %.10f\n", nilai_akar2)
}