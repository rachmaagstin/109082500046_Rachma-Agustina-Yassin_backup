package main

import "fmt"

const NMAX = 1000
type arrFloat [NMAX]float64

func main() {
	var N int
	var data arrFloat

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}

	min := data[0]
	max := data[0]

	for i := 1; i < N; i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}
	fmt.Println(min, max)
}