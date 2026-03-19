# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Rachma Agustina Yassin - 109082500046</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main

import "fmt"

func main() {
var (
	satu, dua, tiga string
	temp string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin/blob/main/modul2/output/output-no1.png)
Program tersebut digunakan untuk menerima tiga input string dari pengguna, kemudian menampilkan urutan nilai sebelum dan sesudah dilakukan pertukaran. Nilai dari variabel satu, dua, dan tiga ditukar menggunakan variabel sementara temp. Proses pertukaran nilai variabel dengan cara memindahkan nilai variabel satu ke variabel temp, kemudian nilai variabel dua dipindahkan ke variabel satu, nilai variabel tiga dipindahkan ke variabel dua, dan nilai variabel yang disimpan di variabel temp dipindahkan ke variabel tiga. Proses ini menghasilkan perubahan urutan secara rotasi, sehingga nilai yang awalnya pertama menjadi terakhir. Program kemudian menampilkan hasil akhir setelah pertukaran nilai dilakukan.

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah "merah", "kuning", "hijau", dan "ungu" selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
#### soal2.go

```go
package main

import "fmt"

func main() {
	var col1, col2, col3, col4 string
	var berhasil bool

	berhasil = true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&col1, &col2, &col3, &col4)

		if !(col1 == "merah" && col2 == "kuning" && col3 == "hijau" && col4 == "ungu") {
			berhasil = false
		}
	}
	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin/blob/main/modul2/output/output-no2.png)
Program tersebut digunakan untuk memeriksa apakah urutan warna pada empat gelas reaksi sesuai dengan ketentuan yaitu merah, kuning, hijau, dan ungu. Program meminta pengguna memasukkan empat warna pada setiap percobaan yang dilakukan sebanyak lima kali menggunakan perulangan. Setiap input percobaan yang dimasukkan kemudian dibandingkan dengan urutan warna yang benar. Jika semua percobaan sesuai, maka program menampilkan BERHASIL: true, tetapi jika ada satu percobaan yang tidak sesuai maka hasilnya BERHASIL: false.


### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
#### soal3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rachmaagstin/109082500046_Rachma-Agustina-Yassin/blob/main/modul2/output/output-no3.png)
Program tersebut digunakan untuk menentukan biaya pengiriman paket berdasarkan beratnya. Dengan cara memecah berat total paket menjadi dua bagian utama yaitu ada berat kilogram utuh dan sisa gram yang kurang dari 1 kilogram. Awalnya, cara kerja program akan meminta untuk memasukkan berat total paket, semuanya dalam satuan gram. Dalam proses penentuan biaya pokok (kilogram) dengan menghitung berapa biaya dasar untuk berat utuh dalam kilogram. Misalnya, ketika dimasukkan berat total paket 8500 gram, maka kita akan mendapatkan berat utuh dalam kilogram sebesar 8 kilogram, ibaratkan setiap kilogram kita dikenakan biaya Rp 10.000. Jadi, jika paket 8 kilogram, biaya dasarnya langsung dihitung menjadi Rp 80.000. Dengan berat sisa 500 gram, makan kita melakukan penentuan biaya tambahan (gram sisa). Biaya tambahan ini untuk menghitung sisa gram yang kurang dari 1 kilogram. Dengan catatan jika ternyata berat utuh paket mencapai 10 kilogram atau lebih, secara otomatis akan diberikan diskon penuh pada sisa gramnya. Artinya, sisa gram itu dianggap gratis, sehingga biaya tambahannya menjadi nol. Tetapi jika paket kurang dari 10 kilogram, maka akan dihitung jumlah sisa gramnya. Nah apabila sisa gram nya sedikit (di bawah 500 gram) biayanya akan lebih mahal, yaitu sebesar Rp. 15 per gram. Sedangkan jika sisa gram nya banyak (500 gram ke atas) biayanya akan lebih murah, hanya Rp 5 per gram. Karena berat sisa yang dijelasin sebelumnya 500 gram maka
kita dikenakan biaya sebesar Rp. 2.500. Sehingga total harga menjadi Rp. 82.500.