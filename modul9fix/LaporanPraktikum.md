# <h1 align="center">Laporan Praktikum Modul 9 - ARRAY </h1>
<p align="center">Rachma Agustina Yassin - 109082500046</p>

## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
#### no1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul9fix/output/no1.png)
Program tersebut digunakan untuk menentukan posisi sebuah titik terhadap dua lingkaran berdasarkan koordinat yang diinputkan nantinya. Program memiliki fungsi jarak untuk menghitung jarak antara dua titik menggunakan akar dari selisih kuadrat koordinat, serta fungsi diDalam untuk mengecek apakah suatu titik berada di dalam lingkaran dengan cara membandingkan jarak titik ke pusat lingkaran dengan jari-jari lingkaran. Pada fungsi main, program membaca input berupa pusat dan radius dua lingkaran, serta koordinat sebuah titik, kemudian mengecek apakah titik tersebut berada di dalam lingkaran pertama, lingkaran kedua, keduanya, atau di luar keduanya. Program bekerja dengan menggunakan satu struc titikLingkaran yang menyimpan koordinat (x, y) dan radius (r), di mana untuk titik hanya digunakan x dan y. Pertama, program menerima input pusat dan radius dua lingkaran (l1 dan l2), serta satu titik (t). Kemudian fungsi jarak(p, q) menghitung jarak antara dua titik menggunakan rumus jarak Euclidean (akar dari selisih kuadrat koordinat), dan fungsi diDalam(c, p) mengecek apakah jarak antara pusat lingkaran dan titik lebih kecil atau sama dengan radius lingkaran (artinya titik berada di dalam lingkaran). Hasil pengecekan disimpan dalam variabel boolean dalam1 dan dalam2, lalu digunakan dalam percabangan if-else untuk menentukan dan mencetak posisi akhir titik sesuai kondisi yang terpenuhi. 



### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut: a. Menampilkan keseluruhan isi dari array. b. Menampilkan elemen-elemen array dengan indeks ganjil saja. c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap). d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna. e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil. f. Menampilkan rata-rata dari bilangan yang ada di dalam array. g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut. h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.
#### no2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul9fix/output/no2.png)
Program tersebut digunakan untuk mengolah data angka dengan menggunakan array dengan cara menerima input data dari pengguna, lalu mengolahnya langkah demi langkah hingga menghasilkan berbagai informasi. Pertama, program meminta pengguna memasukkan jumlah elemen, kemudian membuat array dengan ukuran tersebut. Setelah itu, program mengisi array dengan angka-angka yang dimasukkan satu per satu melalui perulangan. Jika semua data sudah masuk, program langsung menampilkan seluruh isi array. Selanjutnya, program memproses data dengan menampilkan elemen berdasarkan indeks, yaitu indeks ganjil dan genap, menggunakan operasi modulus untuk mengecek posisi indeks. Setelah itu, program meminta nilai x dan menampilkan elemen yang berada pada indeks kelipatan x. Berikutnya, program dapat menghapus elemen pada indeks tertentu yang dimasukkan pengguna dengan cara menggabungkan bagian array sebelum dan sesudah indeks tersebut. Setelah proses penghapusan, program menghitung rata-rata dengan menjumlahkan semua elemen yang tersisa lalu membaginya dengan jumlah data. Kemudian, program menghitung standar deviasi dengan mencari selisih setiap elemen terhadap rata-rata, mengkuadratkannya, menjumlahkan hasilnya, lalu membaginya dengan jumlah data dan diakhiri dengan pengambilan akar kuadrat. Terakhir, program meminta pengguna memasukkan sebuah angka yang ingin dicari, lalu menghitung berapa kali angka tersebut muncul dalam array dengan cara memeriksa setiap elemen satu per satu.
 


### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
#### no3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	hasil := []string{}
	pertandingan := 1

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
		pertandingan++
	}

	for i := 0; i < len(hasil); i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul9fix/output/no3.png)
Program ini digunakan untuk mencatat dan menampilkan hasil pertandingan antara dua klub secara berulang. Cara kerjanya dimulai dengan pengguna memasukkan nama klub A dan klub B, kemudian program akan meminta skor dari kedua klub untuk setiap pertandingan. Program menggunakan perulangan (loop) agar proses input skor bisa dilakukan berkali-kali tanpa perlu menjalankan ulang program. Di dalam perulangan tersebut, terdapat percabangan (if-else) yang berfungsi membandingkan skor untuk menentukan hasil pertandingan, yaitu jika skor klub A lebih besar maka klub A menang, jika skor klub B lebih besar maka klub B menang, dan jika sama maka hasilnya draw. Selain itu, terdapat logika penghentian, yaitu ketika salah satu skor bernilai negatif, maka perulangan akan berhenti. Setiap hasil pertandingan kemudian disimpan ke dalam array menggunakan proses append, append yaitu menambahkan data baru ke dalam array tanpa harus membuat array baru secara manual. Setelah semua pertandingan selesai, program menggunakan perulangan lagi untuk menampilkan seluruh hasil pertandingan satu per satu sesuai urutan, lalu menutup dengan pesan bahwa pertandingan telah selesai.



### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.
#### no4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch string
	*n = 0

	for {
		fmt.Scan(&ch) 

		if ch == "." || *n >= NMAX {
			break
		}

		t[*n] = rune(ch[0])
		(*n)++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](ttps://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul9fix/output/no4.png)
Program ini digunakan untuk membalik atau reverse urutan huruf yang dimasukkan oleh pengguna. Cara kerjanya dimulai saat pengguna memasukkan teks satu per satu (huruf per huruf), lalu program menyimpannya ke dalam sebuah array bertipe rune. Proses input dilakukan terus-menerus hingga pengguna memasukkan tanda titik (.) sebagai tanda berhenti atau jika kapasitas array sudah penuh. Setelah semua huruf disimpan, program akan memanggil fungsi untuk membalik isi array, yaitu dengan menukar posisi elemen dari depan dengan elemen dari belakang secara berpasangan hingga mencapai tengah array. Setelah proses pembalikan selesai, program kemudian menampilkan hasilnya dengan mencetak seluruh isi array yang sudah dalam kondisi terbalik. 



### Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.
#### no4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for *n < NMAX {
		fmt.Scanf("%c", &ch)

		if ch == '\n' || ch == '\r' {
			break
		}

		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			t[*n] = ch
			(*n)++
		}
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var temp tabel

	for i := 0; i < n; i++ {
		temp[i] = t[i]
	}

	balikanArray(&temp, n)

	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}

	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Teks : ")
	isiArray(&tab, &n)

	fmt.Print("Palindrom ? ")
	fmt.Println(palindrom(tab, n))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin_backup/blob/main/modul9fix/output/no4modif.png)
Program hasil modifikasi ini digunakan untuk mengecek apakah suatu teks merupakan palindrom, yaitu teks yang dibaca sama dari depan dan belakang, dengan hanya memperhatikan huruf saja. Cara kerjanya dimulai ketika pengguna memasukkan teks, lalu program membaca input huruf per huruf menggunakan fungsi isiArray hingga menemui enter atau kapasitas penuh. Dalam proses ini, hanya huruf (A–Z dan a–z) yang disimpan ke dalam array, sementara spasi dan karakter lain diabaikan. Setelah itu, fungsi palindrom dipanggil, yang terlebih dahulu menyalin isi array ke array sementara, kemudian membalik urutannya menggunakan fungsi balikanArray dengan cara menukar elemen depan dan belakang. Selanjutnya, program membandingkan array asli dengan array yang sudah dibalik, jika semua elemen sama maka hasilnya true (palindrom), jika tidak maka false. Terakhir, hasil tersebut ditampilkan ke layar. 