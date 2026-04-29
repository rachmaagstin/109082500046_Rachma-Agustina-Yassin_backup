# <h1 align="center">Laporan Praktikum Modul 10 - PENCARIAN NILAI MAX/MIN </h1>
<p align="center">Rachma Agustina Yassin - 109082500046</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
#### no1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul10/output/no1.png)
Program tersebut digunakan untuk menganalisis berat anak kelinci yang akan dijual, dengan cara menentukan berat terkecil dan terbesar dari sejumlah data berat kelinci yang dimasukkan. Cara kerjanya program dimulai dengan membaca input sebuah bilangan bulat (N) sebagai jumlah data, kemudian program menginput (N) nilai berat dan menyimpannya ke dalam array. Setelah semua data tersimpan, program menentukan nilai minimum dan maksimum menggunakan elemen pertama array, lalu melakukan perulangan untuk membandingkan setiap elemen berikutnya. Jika ditemukan nilai yang lebih kecil dari minimum maka nilai minimum diperbarui, dan jika ditemukan nilai yang lebih besar dari maksimum maka nilai maksimum diperbarui. Setelah seluruh data diperiksa, program menampilkan dua nilai yaitu berat terkecil dan terbesar sebagai hasil akhirnya.


### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### no2.go

```go
package main

import "fmt"

const NMAX = 1000
type arrFloat [NMAX]float64

func main() {
	var x, y int
	var ikan arrFloat

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var jumlahWadah int
	var totalWadah arrFloat
	var idx int = 0

	for i := 0; i < x; i++ {
		totalWadah[idx] += ikan[i]

		if (i+1)%y == 0 {
			idx++
		}
	}

	if x%y != 0 {
		jumlahWadah = x/y + 1
	} else {
		jumlahWadah = x / y
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(totalWadah[i], " ")
	}
	fmt.Println()

	var totalSemua float64
	for i := 0; i < jumlahWadah; i++ {
		totalSemua += totalWadah[i]
	}

	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Println(rataRata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul10/output/no2.png)
Program tersebut digunakan untuk mengelompokkan data, misalnya berat ikan atau jumlah ikan ke dalam beberapa wadah, kemudian menghitung total setiap wadah dan rata-rata keseluruhan. Cara kerjanya program dimulai dari menginputkan dua angka, yaitu: x (jumlah data ikan) dan y (kapasitas tiap wadah). Setelah itu, program membaca nilai x data ikan dan menyimpannya dalam array. Selanjutnya, program menjumlahkan nilai ikan ke dalam array totalWadah berdasarkan kelompok berukuran y. Setiap kali jumlah nilai yang telah dimasukkan ke wadah mencapai y, indeks wadah akan berpindah ke berikutnya. Jika jumlah data tidak habis dibagi y, maka wadah terakhir tetap dihitung. Setelah itu, program mencetak total isi setiap wadah. Terakhir, semua total wadah dijumlahkan lalu dibagi dengan jumlah wadah untuk mendapatkan rata-rata isi wadah. Inti dari kegunaan program ini berfungsi untuk membagi data ke dalam kelompok-kelompok dan menganalisis total serta rata-ratanya. 


### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### no3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
	fmt.Printf("Berat balita minimum: %.2f kg\n", *bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", *bMax)
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	var min, max float64

	hitungMinMax(data, n, &min, &max)
	rata := rerata(data, n)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul10/output/no3.png)
Program tersebut digunakan untuk mengolah data berat balita, khususnya untuk mencari berat minimum, maksimum, dan rata-rata. Cara kerja program dimulai dengan meminta jumlah data (n), lalu pengguna memasukkan berat balita satu per satu ke dalam array. Fungsi hitungMinMax kemudian digunakan untuk mencari nilai terkecil dan terbesar. Fungsi ini bekerja dengan membuat nilai minimum dan maksimum dari elemen pertama, lalu membandingkannya dengan elemen lain secara berulang untuk menemukan nilai paling kecil dan paling besar. Hasilnya dicetak ke layar. Setelah itu, fungsi rerata menghitung rata-rata dengan menjumlahkan semua data berat dan membaginya dengan jumlah data. Hasil rata-rata juga ditampilkan. Secara keseluruhan, program ini berfungsi untuk menganalisis data numerik sederhana dengan mencari nilai minimum, maksimum, dan rata-rata.