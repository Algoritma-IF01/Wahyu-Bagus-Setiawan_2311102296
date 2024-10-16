package main
import "fmt"
//tahun kabisat

func main (){
	var tahun int

	fmt.Print("Masukkan Tahun : ")
	fmt.Scanln (&tahun)

	if tahun%4 == 0 {
		fmt.Println("Tahun ini Tahun Kabisat ")
	}else {
		fmt.Println("Tahun Ini bkn Tahun Kabisat")
	}
}