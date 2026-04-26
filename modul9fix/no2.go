package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi dari array.
	fmt.Println("\nSemua elemen:", arr)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
	fmt.Print("Indeks ganjil: ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.
	var idx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Array setelah dihapus:", arr)

	// f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata:", mean)

	// g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.
	var total float64
	for i := 0; i < len(arr); i++ {
		total += math.Pow(float64(arr[i])-mean, 2)
	}
	stdDev := math.Sqrt(total / float64(len(arr)))
	fmt.Println("Standar deviasi:", stdDev)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.
	var cari int
	fmt.Print("Masukkan angka yang dicari: ")
	fmt.Scan(&cari)

	freq := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}