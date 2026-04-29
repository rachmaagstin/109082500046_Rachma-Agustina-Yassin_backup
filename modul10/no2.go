package main

import "fmt"

const NMAX = 1000
type arrFloat [NMAX]float64

func main() {
	var x, y int
	var ikan arrFloat

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var jumlahWadah int
	var totalWadah arrFloat
	var idx int = 0

	for i := 0; i < x; i++ {
		totalWadah[idx] += ikan[i]

		if (i+1)%y == 0 {
			idx++
		}
	}

	if x%y != 0 {
		jumlahWadah = x/y + 1
	} else {
		jumlahWadah = x / y
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(totalWadah[i], " ")
	}
	fmt.Println()

	var totalSemua float64
	for i := 0; i < jumlahWadah; i++ {
		totalSemua += totalWadah[i]
	}

	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Println(rataRata)
}