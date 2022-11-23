package main

import "fmt"

func main() {

	type NoEktp string
	type merried bool
	type Nama string
	type TanggalLahir string
	type TempatLahir string

	var NoEktpomte NoEktp = "346326473264927"
	var merriedStatus merried = true
	var DetailNama Nama = "Khoirul Fikri"
	var DetailTglLahir TanggalLahir = "20-oktober-1995"
	var DetailTempatLahir TempatLahir = "Yogyakarta"

	fmt.Println(NoEktpomte)
	fmt.Println(merriedStatus)
	fmt.Println(DetailNama)
	fmt.Println(DetailTglLahir)
	fmt.Println(DetailTempatLahir)

}
