package main
import "fmt"
func main (){
	var greetings = "Selamat Pagi"
	var a,b int

	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n" , a , b , a+b)
}