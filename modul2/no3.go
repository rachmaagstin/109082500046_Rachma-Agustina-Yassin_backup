package main

import "fmt"

func main() {
	var berat_total, biaya_sisa, total int
	var kilogram = berat_total / 1000
	var biayaPerKg = kilogram * 10000
	var sisagram = berat_total % 1000

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&berat_total)

	kilogram = berat_total / 1000
	biayaPerKg = kilogram * 10000
	sisagram = berat_total % 1000
	biaya_sisa = 0

	if sisagram >= 500 {
	biaya_sisa = sisagram * 5
	} else {
	biaya_sisa = sisagram * 15
	}

	total = biayaPerKg
	if kilogram <= 10 {
	total += biaya_sisa
	}
	
	fmt.Printf("Total berat: %d kg dan %d gram\n", kilogram, sisagram)
	fmt.Printf("Total biaya kirim paket: Rp. %d + Rp. %d\n", biayaPerKg, biaya_sisa)
	fmt.Printf("Total harga: Rp. %d\n", total)
}