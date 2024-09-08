package main

import (
	"fmt"
)

func main() {

	// Constant Variabel Tidak dapat diubah setelah diberi nilai pertama kali
	// Cara Pembuatan Constant mirip dengan Variabel dan wajib lansung menginisialisasikan datanya

	const (
		nama = "Tobiichi Origami"
		umur = ""
	)

	fmt.Println(nama, umur)

	// konversi tipe data numeric (1)
	// jika tipe data overflow = kembali ke belakang
	var (
		nilai1 int64 = 98500
		nilai2 int32 = int32(nilai1)
		nilai3 int16 = int16(nilai2)
	)

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai3)
	fmt.Println()

	// konversi tipe data string (2)
	var (
		name          = "Origami"
		e       uint8 = name[6]   // e adalah byte
		estring       = string(e) // konversi byte ke string
	)

	fmt.Println(e)       // Akan menampilkan nilai byte dari karakter 'O'
	fmt.Println(estring) // Akan menampilkan karakter 'O'
	fmt.Println()

	// type Declarations
	// Type Declarations membuat ulang data baru dari tipe yang sudah ada
	// Type Declarations biasanya digunakan untuk membuat alias terhadap data yang sudah ada,dengan tujuan agar lebih mudah dimengerti.

	type Umur int
	var UmurOrigami Umur = 18
	var tahunlalu int = 17
	var dulu Umur = Umur(tahunlalu)
	fmt.Println(UmurOrigami)
	fmt.Println(Umur(8))
	fmt.Println(Umur(dulu))
	fmt.Println()

	// Operasi Matematika

	var (
		a = 10
		b = 20
		c = 30
		d = 40
		f = a + b - c/d
	)
	fmt.Println(f)
	fmt.Println()

	// Augmented Assigment
	var (
		i = 10
	)
	i += 10
	fmt.Println(i)

	// Unary Operator
	var (
		G = 1
	)
	G++
	fmt.Println(G)
	G--
	fmt.Println(G)
	fmt.Println()

	// operasi perbandingan
	// operasi perbandingan adalah operasi yang membandikan dua buah data
	// biasa menghasilkan boolean (True OR False)
	var (
		number1 int = 1
		number2 int = 2

		anime1 = "DateALive"
		anime2 = "DateALive"

		waifu1 = "Tobiichi Origami"
		waifu2 = "Itsuka Chiyogami"
	)

	var result2 bool = anime1 == anime2
	var result bool = number1 >= number2
	var result3 bool = waifu1 != waifu2

	fmt.Println("1 lebih besar dari 2 = ", result)
	fmt.Println("Judul Animenya Apakah sama = ", result2)
	fmt.Println("Kedua Waifuku Apakah beda = ", result3)
	fmt.Println()

	// Operasi Boolean
	// && = dan
	// || = OR
	// ! = kebalikan

	var (
		grade_A = 90
		exam    = 100

		nilaisiswa1 bool = 100 >= grade_A
		examsiswa1  bool = 100 >= exam

		nilaisiswa2 bool = 100 > grade_A
		absensiswa2 bool = 80 > exam

		// nilaisiswa3 bool = 100 > grade_A
		// absensiswa3 bool = 1 > absen

		pass1 = nilaisiswa1 && examsiswa1
		pass2 = nilaisiswa2 || absensiswa2
	)

	fmt.Println("Apakah Siswa 1 Lulus? = ", pass1)
	fmt.Println("Apakah Siswa 2 Lulus? = ", pass2)

	/* DAY 1

	// about integer and data floating
	// complex64 float32 int8 unint8

	fmt.Println("satu = ", 1)
	fmt.Println(`dua = `, 2)
	fmt.Println(`Tiga koma lima = `, 3.5)

	fmt.Println()
	// boolean
	fmt.Println("true = ", true)
	fmt.Println("false = ", false)

	// string
	fmt.Println("")
	fmt.Println("String")

	fmt.Println("")

	// function for String
	fmt.Println(len("wkwkwkwk")) //menghitung jumlah string
	fmt.Println("nani"[1])       // menghitung ASCII string

	// Variable
	var name string
	name = "Chiyogami"
	fmt.Println(name)

	name = "Tobiichi Origami"
	fmt.Println(name)

	Country := "Indonesia +62" //menambahkan : di := akan membuat variable gak bisa diulang contoh nama := dan nama := akan error
	fmt.Println(Country)

	Country = "japan +81"
	fmt.Println(Country)

	// Deklarasi Multiple Variable
	var (
		Nama = "Tobiichi Origammi"
		Umur = 18
		// wajib deklarasikan dan gunakan variablenya atau akan error
	)

	fmt.Println(Nama)
	fmt.Println(Umur)

	*/
}
