# <h1 align="center">Laporan Praktikum Modul 14 - SORTING </h1>
<p align="center">Rachma Agustina Yassin - 109082500046</p>

## Unguided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing-masing daerah.
#### no1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul14/no1.go)
Program tersebut digunakan untuk membaca beberapa kumpulan bilangan, kemudian mengurutkan setiap kumpulan data dari kecil ke besar menggunakan cara selection sort dan menampilkan hasil pengurutannya. Program dimulai dengan membaca jumlah input yang disimpan pada variabel m, kemudian melakukan perulangan sebanyak jumlah input tersebut. Pada setiap perulangan, program membaca jumlah data n dan menyiapkan array rumah berukuran 100 elemen untuk menyimpan bilangan yang diinput pengguna. Selanjutnya program membaca setiap bilangan dan menyimpannya langsung ke dalam array j. Setelah seluruh data tersimpan, fungsi rumahKerabat dipanggil dengan alamat array &rumah dan jumlah data n. Fungsi tersebut menentukan indeks nilai terkecil sementara pada variabel idxMin, lalu membandingkannya dengan seluruh elemen setelahnya untuk mencari nilai yang lebih kecil. Jika ditemukan nilai yang lebih kecil, indeks minimum diperbarui. Setelah proses pencarian selesai, elemen pada posisi sekarang ditukar dengan elemen terkecil yang ditemukan. Proses ini diulang hingga seluruh data terurut dari kecil ke besar. Setelah pengurutan selesai, program mencetak seluruh isi array yang sudah terurut.


### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Format Masukan masih persis sama seperti sebelumnya. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.
#### no2.go

```go
package main

import "fmt"

func kerabatDekatMin(data *[100]int, n int) {
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

func kerabatDekatMax(data *[100]int, n int) {
	for i := 0; i < n-1; i++ {
		max := i

		for j := i + 1; j < n; j++ {
			if data[j] > data[max] {
				max = j
			}
		}

		data[i], data[max] = data[max], data[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for ; n > 0; n-- {
		var m int
		fmt.Scan(&m)

		var ganjil [100]int
		var genap [100]int

		jumlahGanjil := 0
		jumlahGenap := 0

		for i := 0; i < m; i++ {
			var x int
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[jumlahGanjil] = x
				jumlahGanjil++
			} else {
				genap[jumlahGenap] = x
				jumlahGenap++
			}
		}

		kerabatDekatMin(&ganjil, jumlahGanjil)
		kerabatDekatMax(&genap, jumlahGenap)

		for i := 0; i < jumlahGanjil; i++ {
			fmt.Print(ganjil[i], " ")
		}

		for i := 0; i < jumlahGenap; i++ {
			fmt.Print(genap[i], " ")
		}

		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul14/no2.go)
Program tersebut digunakan untuk memisahkan bilangan ganjil dan genap dari beberapa kumpulan data, kemudian mengurutkan bilangan ganjil dari kecil ke besar dan bilangan genap dari besar ke kecil menggunakan cara selection sort, lalu menampilkan hasilnya. Program dimulai dengan membaca jumlah input n, kemudian pada setiap perulangan membaca jumlah data m dan setiap bilangan yang dimasukkan pengguna. Jika bilangan ganjil maka disimpan ke array ganjil, sedangkan bilangan genap disimpan ke array genap. Setelah semua data tersimpan, fungsi kerabatDekatMin dipanggil untuk mengurutkan array ganjil dengan cara mencari nilai terkecil lalu menukarnya ke posisi depan, sedangkan fungsi kerabatDekatMax digunakan untuk mengurutkan array genap dengan cara mencari nilai terbesar lalu menukarnya ke posisi depan. Kedua fungsi menggunakan alamat array agar perubahan langsung memengaruhi array asli tanpa menyalin data. Setelah proses pengurutan selesai, program mencetak seluruh bilangan ganjil yang sudah terurut terlebih dahulu, kemudian diikuti bilangan genap yang juga sudah terurut. Setelah mencoba input pada soal, hasil output berbeda dengan yang ada di soal, mungkin output pada soal 2 terdapat kekeliruan.


### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. Keluaran adalah median yang diminta, satu data per baris.
#### no3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul14/no3.go)
Program tersebut digunakan untuk membaca bilangan yang diinputkan lalu mencetak nilai median setiap kali muncul input bernilai 0. Program dimulai dengan membaca input satu per satu dan menyimpannya ke dalam array angka selama nilainya bukan 0 dan bukan -5313. Nilai -5313 digunakan sebagai penanda akhir input sehingga program berhenti. Ketika program membaca angka 0, seluruh data yang sudah tersimpan disalin ke array sementara yaitu temp agar data tidak berubah, kemudian fungsi median dipanggil untuk mengurutkan data menggunakan cara selection sort dengan cara mencari nilai terkecil pada bagian array yang belum terurut lalu menukarnya ke posisi depan hingga seluruh data terurut dari terkecil ke terbesar. Setelah data terurut, program menghitung median, yaitu nilai tengah jika jumlah data ganjil, atau rata-rata dua nilai tengah jika jumlah data genap. Hasil median kemudian dicetak ke layar setiap kali ditemukan input 0.