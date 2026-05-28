package main

import "fmt"

func kerabatDekatMin(data *[100]int, n int) {
	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; j < n; j++ {
			if data[j] < data[min] {
				min = j
			}
		}

		data[i], data[min] = data[min], data[i]
	}
}

func kerabatDekatMax(data *[100]int, n int) {
	for i := 0; i < n-1; i++ {
		max := i

		for j := i + 1; j < n; j++ {
			if data[j] > data[max] {
				max = j
			}
		}

		data[i], data[max] = data[max], data[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for ; n > 0; n-- {
		var m int
		fmt.Scan(&m)

		var ganjil [100]int
		var genap [100]int

		jumlahGanjil := 0
		jumlahGenap := 0

		for i := 0; i < m; i++ {
			var x int
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[jumlahGanjil] = x
				jumlahGanjil++
			} else {
				genap[jumlahGenap] = x
				jumlahGenap++
			}
		}

		kerabatDekatMin(&ganjil, jumlahGanjil)
		kerabatDekatMax(&genap, jumlahGenap)

		for i := 0; i < jumlahGanjil; i++ {
			fmt.Print(ganjil[i], " ")
		}

		for i := 0; i < jumlahGenap; i++ {
			fmt.Print(genap[i], " ")
		}

		fmt.Println()
	}
}