package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= int(i)
	}
}

func permutasi(n, r int, hasil *int) {
	var faktorial_n, faktorial_nMINr int
	faktorial(n, &faktorial_n)
	faktorial(n-r, &faktorial_nMINr)
	*hasil = faktorial_n / faktorial_nMINr
}

func kombinasi(n, r int, hasil *int) {
	var faktorial_n, faktorial_r, faktorial_nMINr int
	faktorial(n, &faktorial_n)
	faktorial(r, &faktorial_r)
	faktorial(n-r, &faktorial_nMINr)
	*hasil = faktorial_n / (faktorial_r * faktorial_nMINr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	var permutasi1, kombinasi1, permutasi2, kombinasi2 int

	permutasi(a, c, &permutasi1)
	kombinasi(a, c, &kombinasi1)
	permutasi(b, d, &permutasi2)
	kombinasi(b, d, &kombinasi2)

	fmt.Println(permutasi1, kombinasi1)
	fmt.Println(permutasi2, kombinasi2)
}