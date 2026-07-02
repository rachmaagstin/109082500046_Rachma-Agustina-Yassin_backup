package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	var data []int
	var x int

	for {
		fmt.Scan(&x)
		if x == -5313 {
			break
		}
		if x == 0 {
			temp := make([]int, len(data))
			copy(temp, data)
			insertionSort(temp)

			n := len(temp)

			if n%2 == 1 {
				fmt.Println(temp[n/2])
			} else {
				median := (temp[n/2-1] + temp[n/2]) / 2
				fmt.Println(median)
			}

		} else {
			data = append(data, x)
		}
	}
}