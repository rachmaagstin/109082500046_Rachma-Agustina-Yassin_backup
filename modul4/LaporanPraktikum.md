# <h1 align="center">Laporan Praktikum Modul 4 - PROSEDUR </h1>
<p align="center">Rachma Agustina Yassin - 109082500046</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p). Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n!/(n−r)!, sedangkan C(n, r) = n!/r!(n−r)!
#### no1.go

```go
package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= int(i)
	}
}

func permutasi(n, r int, hasil *int) {
	var faktorial_n, faktorial_nMINr int
	faktorial(n, &faktorial_n)
	faktorial(n-r, &faktorial_nMINr)
	*hasil = faktorial_n / faktorial_nMINr
}

func kombinasi(n, r int, hasil *int) {
	var faktorial_n, faktorial_r, faktorial_nMINr int
	faktorial(n, &faktorial_n)
	faktorial(r, &faktorial_r)
	faktorial(n-r, &faktorial_nMINr)
	*hasil = faktorial_n / (faktorial_r * faktorial_nMINr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	var permutasi1, kombinasi1, permutasi2, kombinasi2 int

	permutasi(a, c, &permutasi1)
	kombinasi(a, c, &kombinasi1)
	permutasi(b, d, &permutasi2)
	kombinasi(b, d, &kombinasi2)

	fmt.Println(permutasi1, kombinasi1)
	fmt.Println(permutasi2, kombinasi2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul3/output/no1.png)
Program tersebut digunakan untuk menghitung nilai permutasi dan kombinasi dari dua pasang bilangan yang diinputkan dengan menggunakan prosedur, yaitu (a, c) dan (b, d). Dalam program tersebut menggunakan beberapa prosedur(func) untuk memecah setiap proses perhitungan agar lebih mudah untuk digunakan kembali. Pada (func) pertama digunakan untuk menghitung faktorial yang lalu hasil nya disimpan. Selanjutnya, pada (func) kedua yaitu permutasi memanggil (func) faktorial untuk menghitung banyaknya susunan dengan memperhatikan urutannya. Selanjutnya, pada (func) ketiga yaitu kombinasi memanggil (func) faktorial juga untuk menghitung banyaknya susunan tetapi tidak memperhatikan urutannya. Kemudian, pada (func) main diminta untuk memasukkan empat bilangan (dua pasang bilangan), lalu program akan menghitung menggunakan (func) faktorial, permutasi dan kombinasi tanpa harus ditulis ulang perhitungannya secara berkali-kali, karena sudah terpecah dan teratur sehingga mudah untuk diproses.


### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer). Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
#### no2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul3/output/no2.png)
Program tersebut digunakan untuk menghitung skor pada perlombaan dengan menggunakan prosedur. Pada (func) hitungSkor memiliki dua tempat penyimpanan, yaitu soal dan skor yang digunakan untuk menghitung berapa soal yang telah dikerjakan dan juga untuk mengetahui total waktu pengerjaan soalnya. Di dalam (func) hitungSkor, program membaca waktu pengerjaan untuk 8 soal satu per satu. Jika waktu suatu soal kurang dari atau sama dengan 300 detik, maka soal tersebut dianggap berhasil diselesaikan, sehingga jumlah soal bertambah satu dan waktunya ditambahkan ke total skor. Selanjutnya, pada (func) main digunakan untuk membaca nama peserta yang diinputkan secara berulang, jika telah diinputkan "Selesai" maka program akan berhenti meminta input. Jadi, setiap kali ada nama peserta yang diinputkan dan terbaca maka program akan memanggil (func) hitungSkor untuk menghitung berapa banyak pengerjaan soal pesertanya, lalu hasilnya akan dibandingkan dengan peserta lain, dan peserta yang mendapat nilai tertinggi sebagai pemenangnya. Kemudian, outputnya akan ada nama pemenang serta berapa banyak soal yang telah ia kerjakan dengan total waktunya.