package main

import "fmt"

func rumahKerabat(data *[100]int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i

		for j := i + 1; j < n; j++ {
			if data[j] < data[idxMin] {
				idxMin = j
			}
		}

		data[i], data[idxMin] = data[idxMin], data[i]
	}
}

func main() {
	var m int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		var n int
		fmt.Scan(&n)

		var rumah [100]int

		for j := 0; j < n; j++ {
			fmt.Scan(&rumah[j])
		}

		rumahKerabat(&rumah, n)

		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}

			fmt.Print(rumah[j])
		}

		fmt.Println()
	}
}