package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	isNumber := func(s string) bool {
		_, err := strconv.Atoi(s)
		return err == nil
	}
	fmt.Print("Masukkan Input String Pertama: ")
	fmt.Scanln(&satu)
	if isNumber(satu) {
		fmt.Println("Error: Input Anda Harus Huruf/String bkn Angka/integer")
		return
	}
	fmt.Print("Masukkan Input String Kedua: ")
	fmt.Scanln(&dua)
	if isNumber(dua) {
		fmt.Println("Error: Input Anda Harus Huruf/String bkn Angka/integer")
		return
	}
	fmt.Print("Masukkan Input String Ketiga: ")
	fmt.Scanln(&tiga)
	if isNumber(tiga) {
		fmt.Println("Error: Input Anda Harus Huruf/String bkn Angka/integer")
		return
	}
	fmt.Println("Output awal = " + satu + "" + dua + "" + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output Akhir = " + satu + "" + dua + "" + tiga)
}
