package main
import (
	"fmt"
)

// Procedure untuk ngisi si array
func soalSatuDua(arr *[256]int, n int) {
	fmt.Println("Masukkan", n, "bilangan bulat:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
}

// Procedure untuk balikin isi arraynya
func soalTiga(arr *[256]int, n int) {
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

// Function untuk ngecek arraynya simetris(Palindrom) apa ngga 
func soalEmpat(arr [256]int, n int) bool {
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}
	return true
}

func nama(){
	fmt.Println(" ~ Wahyu Bagus Setiawan ~ ")
	fmt.Println("     -> 2311102296 <-      ")
}
func main() {
	var arr [256]int 
	var n int

	nama()

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	if n > 256 {
		fmt.Println("Jumlah elemen tidak boleh lebih dari 256.")
		return
	}

	//ini buat Design Outputnya jadi enak dilihat
	soalSatuDua(&arr, n)
	fmt.Println("Array yang dimasukkan:", arr[0:n])
	soalTiga(&arr, n)
	fmt.Println("Array setelah dibalik:", arr[0:n])
	if soalEmpat(arr, n) {
		fmt.Println("Array membentuk pola simetris (palindrom).")
	} else {
		fmt.Println("Array tidak membentuk pola simetris (palindrom).")
	}
}
