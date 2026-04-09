# <h1 align="center">Laporan Praktikum Modul 5 - REKURSIF </h1>
<p align="center">Rachma Agustina Yassin - 109082500046</p>

## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut. n 0 1 2 3 4 5 6 7 8 9 10. Sn 0 1 1 2 3 5 8 13 21 34 55.
#### no1.go

```go
package main 

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan n : ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul5/output/no1.png)
Program tersebut digunakan untuk menghitung deret fibonacci dengan menggunakan rekursif. Pada (func) fibonacci terdapat kondisi untuk menentukan nilai (n) itu 0 atau 1, ataupun lebih dari 1. Selanjutnya, pada (func) main terdapat perulangan, dalam perulangan tersebut digunakan untuk mengulang dari 0 sampai (n) yang diinputkan nantinya. Pada setiap perulangan, program akan memanggil (func) fibonacci(i) untuk menghitung nilai fibonacci pada posisi ke-i. Di dalam fungsi tersebut, jika nilai i adalah 0 maka langsung menghasilkan 0, jika 1 menghasilkan 1, dan jika lebih dari 1 maka fungsi akan menghitung nilainya dengan menjumlahkan dua nilai fibonacci sebelumnya dengan cara memanggil nilainya sendiri, yaitu fibonacci(i-1) + fibonacci(i-2). Pada proses pemanggilan fungsi ini terus berulang sampai mencapai kondisi dasar 0 atau 1, lalu hasilnya dikembalikan ke atas dan dijumlahkan. Setelah setiap nilai berhasil dihitung, program akan menampilkannya ke layar. Proses ini akan terus dilakukan sampai semua suku dari 0 hingga (n) selesai ditampilkan, sehingga terbentuk deret fibonaccinya.


### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user. Masukkan 1. Keluaran ("*").
#### no2.go

```go
package main

import "fmt"

func bintang(hasil int) {
	if hasil == 0 {
		return
	}
	bintang(hasil - 1)
	fmt.Print("*")
}

func pola(baris, n int) {
	if baris > n {
		return
	}
	bintang(baris)
	fmt.Println()
	pola(baris+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	pola(1, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul5/output/no2.png)
Program tersebut digunakan untuk membuat pola bintang berbentuk segitiga dengan jumlah baris sesuai yang diinputkan dengan menggunakan rekursif. Pada (func) bintang digunakan untuk mencetak sejumlah bintang dalam satu baris dengan cara memanggil dirinya sendiri sampai jumlah bintang terpenuhi. Setelah satu baris selesai dicetak, (func) pola(baris, n int) akan mengatur perpindahan ke baris berikutnya dan menambah jumlah bintang yang dicetak. Pada (func) main setelah pengguna memasukkan nilai (N), program akan mencetak pola bintang dari 1 hingga (N) baris, di mana setiap baris memiliki jumlah bintang sesuai nomor barisnya. Proses akan berhenti ketika nilai baris melebihi (N). 


### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### no3.go

```go
package main

import "fmt"

func faktor(N, pembagi int) {
	if pembagi > N {
		return
	}

	if N%pembagi == 0 {
		fmt.Printf("%d ", pembagi)
	}

	faktor(N, pembagi+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)
	faktor(N, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul5/output/no3.png)
Program tersebut digunakan untuk mencari jumlah faktor dari suatu bilangan (N) dengan menggunakan rekursif. Pada (func) faktor akan mengecek dengan memeriksa angka satu per satu mulai dari 1 sampai (N) untuk menentukan apakah angka tersebut dapat membagi (N) tanpa sisa. Jika suatu angka dapat membagi (N) habis, maka angka tersebut merupakan faktor dari (N) dan akan dicetak. Pada (func) faktor (N, pembagi) melakukan proses tersebut dengan memanggil dirinya sendiri menggunakan nilai pembagi + 1 untuk memeriksa angka berikutnya. Proses akan terus berjalan sampai nilai pembagi lebih besar dari (N), yang menandakan bahwa seluruh faktor telah diperiksa dan dicetak.