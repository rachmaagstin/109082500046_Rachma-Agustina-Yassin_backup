package main

import (
	"fmt"
	"math"
)

type titikLingkaran struct {
	x, y int
	r    int
}

func jarak(p, q titikLingkaran) float64 {
	return math.Sqrt(
		math.Pow(float64(p.x-q.x), 2) +
			math.Pow(float64(p.y-q.y), 2),
	)
}

func diDalam(c titikLingkaran, p titikLingkaran) bool {
	return jarak(c, p) <= float64(c.r)
}

func main() {
	var l1, l2, t titikLingkaran

	fmt.Scan(&l1.x, &l1.y, &l1.r)
	fmt.Scan(&l2.x, &l2.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	dalam1 := diDalam(l1, t)
	dalam2 := diDalam(l2, t)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}