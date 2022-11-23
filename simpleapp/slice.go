package main

import "fmt"

func main() {
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	//fungsi slice untuk membagi data array
	// len(slice) untuk mendapatkan panjang slice dan cap(slice) digunakan untuk mendapatkan kapasitas
	var slice1 = bulan[4:7]
	var slice2 = bulan[6:8]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
