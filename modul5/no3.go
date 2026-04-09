package main

import "fmt"

func faktor(N, pembagi int) {
	if pembagi > N {
		return
	}

	if N%pembagi == 0 {
		fmt.Printf("%d ", pembagi)
	}

	faktor(N, pembagi+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)
	faktor(N, 1)
}