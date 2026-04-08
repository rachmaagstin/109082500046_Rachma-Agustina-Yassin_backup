package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var max_soal, min_skor, soal, skor int

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > max_soal || (soal == max_soal && skor < min_skor) {
			max_soal = soal
			min_skor = skor
			pemenang = nama
		}
	}
	fmt.Println(pemenang, max_soal, min_skor)
}