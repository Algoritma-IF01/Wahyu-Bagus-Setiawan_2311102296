package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Print("Banyak mahasiswa: ")
	fmt.Scan(&N)

	validCount := 0

	for i := 0; i < N; i++ {
		var voucher string
		fmt.Print("Nomor Seri Voucher Mahasiswa: ")
		fmt.Scan(&voucher)

		length := len(voucher)

		if length == 5 || length == 6 {
			fmt.Println("Vouchernya Valid Karena Tidak Melebihi 5 atau 6 Digit")
		} else {
			fmt.Println("Vouchernya Tidak Valid Karena Melebihi 5 atau 6 Digit")
		}

		// ini buat ngecek valid 2
		if length == 5 {
			if int(voucher[0])*int(voucher[1]) == int(voucher[3])*int(voucher[4]) {
				fmt.Println("Vouchernya Valid Karena Perkalian Digit Pada 2 Digit pertama & 2 digit terakhir adalah sama.")
			}
		}
		if length == 6 {
			if int(voucher[0])*int(voucher[1]) == int(voucher[4])*int(voucher[5]) {
				fmt.Println("Vouchernya Valid Karena Perkalian Digit Pada 2 Digit pertama & 2 Digit Terakhir Sama.")
			}
		} else {
			fmt.Println("Vouchernya Tidak Valid Karena Perkalian Digit Pada 2 Digit pertama & 2 Digit Terakhir Tidak Sama.")
		}

		// ini buat ngecek valid 3
		if length == 5 {
			midDigit := voucher[2]
			midDigitInt, _ := strconv.Atoi(string(midDigit)) // Perbaikan di sini
			if midDigitInt%2 == 0 {
				fmt.Println("Vouchernya Valid Karna Digit yg di Tengah (1 digit untuk panjang nomor seri 5) Genap.")
				validCount++
			} else {
				fmt.Println("Vouchernya Tidak Valid Karna Digit yg di Tengah (1 digit untuk panjang nomor seri 5) Tidak Genap.")
			}
		}
		if length == 6 {
			midDigit := voucher[3]
			midDigitInt, _ := strconv.Atoi(string(midDigit))
			if midDigitInt%2 == 0 {
				fmt.Println("Vouchernya Valid Karna Digit yg di Tengah (3 digit untuk panjang nomor seri 6) Genap.")
				validCount++
			} else {
				fmt.Println("Vouchernya Tidak ValidKarna Digit yg di Tengah (3 digit untuk panjang nomor seri 6) Tidak Genap.")
			}
		}
	}
}
