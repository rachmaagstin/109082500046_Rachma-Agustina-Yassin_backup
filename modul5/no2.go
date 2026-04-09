package main

import "fmt"

func bintang(hasil int) {
	if hasil == 0 {
		return
	}
	bintang(hasil - 1)
	fmt.Print("*")
}

func pola(baris, n int) {
	if baris > n {
		return
	}
	bintang(baris)
	fmt.Println()
	pola(baris+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	pola(1, n)
}