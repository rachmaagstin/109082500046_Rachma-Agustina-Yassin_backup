package main

import "fmt"

func main() {
	var col1, col2, col3, col4 string
	var berhasil bool

	berhasil = true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&col1, &col2, &col3, &col4)

		if !(col1 == "merah" && col2 == "kuning" && col3 == "hijau" && col4 == "ungu") {
			berhasil = false
		}
	}
	fmt.Println("BERHASIL:", berhasil)
}