# <h1 align="center">Laporan Praktikum Modul 3 - FUNGSI </h1>
<p align="center">Rachma Agustina Yassin - 109082500046</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p). Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut! P(n, r) = n!/(n−r)!, sedangkan C(n, r) = n!/r!(n−r)!
#### no1.go

```go
package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com)
Program tersebut digunakan untuk menghitung nilai permutasi (P) dan kombinasi (C) dari dua bilangan yang diinputkan, yaitu (a, c) dan (b, d), dengan syarat n ≥ r. Program diawali dengan fungsi faktorial yang menghitung nilai faktorial suatu bilangan dengan cara memanggil nilai bilangnya sendiri, yaitu n! = n × (n−1)!. Selanjutnya, fungsi permutasi digunakan untuk menghitung banyaknya susunan nilai r dari nilai n menggunakan rumus P(n, r) = n!/(n−r)!, sedangkan fungsi kombinasi menghitung banyaknya cara memilih nilai r dari nilai n tanpa memperhatikan urutannya dengan rumus C(n, r) = n!/(r!(n−r)!). Pada fungsi main, program membaca empat bilangan bulat sebagai input, lalu memanggil fungsi permutasi dan kombinasi untuk masing-masing bilangan dan menampilkan hasilnya dalam dua baris keluaran.


### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x², g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!
#### no2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com)
Program tersebut digunakan untuk menghitung komposisi beberapa fungsi dalam matematika terhadap tiga buah input bilangan, yaitu a, b, dan c. Terdapat tiga fungsi yang didefinisikan : fungsi f(x) untuk menghitung kuadrat suatu bilangan (x²), fungsi g(x) untuk mengurangi bilangan dengan 2 (x−2), dan fungsi h(x) untuk menambahkan 1 pada bilangan (x+1). Di dalam fungsi main, program mulai membaca tiga input, kemudian menghitung dan menampilkan hasil dari tiga komposisi fungsi, yaitu f(g(h(a))), g(h(f(b))), dan h(f(g(c))). Artinya, setiap nilai input diproses secara berurutan oleh beberapa fungsi dengan mengutamakan (fungsi di dalam dijalankan lebih dulu).


### 3. Diberikan tiga buah fungsi matematika yaitu f (x) = x ^ 2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!
#### no3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64
	var in1, in2 bool

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	in1 = didalam(cx1, cy1, r1, x, y)
	in2 = didalam(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com)
Program tersebut digunakan untuk menentukan posisi sebuah titik terhadap dua lingkaran berdasarkan koordinat yang diinputkan nantinya. Program memiliki fungsi jarak untuk menghitung jarak antara dua titik menggunakan akar dari selisih kuadrat koordinat, serta fungsi didalam untuk mengecek apakah suatu titik berada di dalam lingkaran dengan cara membandingkan jarak titik ke pusat lingkaran dengan jari-jari lingkaran. Pada fungsi main, program membaca input berupa pusat dan radius dua lingkaran, serta koordinat sebuah titik, kemudian mengecek apakah titik tersebut berada di dalam lingkaran pertama, lingkaran kedua, keduanya, atau di luar keduanya. Hasil akhir yang ditampilkan berupa kalimat yang menjelaskan posisi titik tersebut.