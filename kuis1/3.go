package main

import "fmt"

func sumKelipatan4(total int) int {
    var num int
    fmt.Scan(&num)

    if num < 0 {
        return total
    }

    if num%4 == 0 {
        total += num
    }

    return sumKelipatan4(total)
}

func main() {
    fmt.Println("Masukkan angka (bilangan negatif untuk berhenti):")
    fmt.Println("Total kelipatan 4:", sumKelipatan4(0))
}
