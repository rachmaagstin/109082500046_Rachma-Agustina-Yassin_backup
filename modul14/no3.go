package main

import "fmt"

func median(data []int, n int) {
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

func main() {
	var x int
	var angka [1000000]int
	n := 0

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			var temp [1000000]int

			for i := 0; i < n; i++ {
				temp[i] = angka[i]
			}

			median(temp[:], n)

			if n%2 == 1 {
				fmt.Println(temp[n/2])
			} else {
				fmt.Println((temp[n/2-1] + temp[n/2]) / 2)
			}

		} else {
			angka[n] = x
			n++
		}
	}
}