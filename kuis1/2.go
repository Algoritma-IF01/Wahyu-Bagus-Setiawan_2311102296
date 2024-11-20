package main

import "fmt"

func hitungBiaya(menu int) int {
	if menu > 50 {
		return 100000
	} else if menu > 3 {
		return 10000 + (menu-3)*2500
	} else {
		return 10000
	}
}

func main() {
	var M int
	fmt.Scan(&M)

	for i := 0; i < M; i++ {
		var menu, orang int
		var sisa bool
		fmt.Scan(&menu, &orang, &sisa)

		biaya := hitungBiaya(menu)

		if sisa {
			biaya *= orang
		}

		fmt.Println(biaya)
	}
}
