package main

import "fmt"

func main() {
	//forr adalah pengulangan

	var jumlah = 1

	for jumlah <= 3 {
		fmt.Println("perulangan ke", jumlah)
		jumlah++
	}

	//for dengan statement

	for nomor := 1; nomor <= 3; nomor++ {
		fmt.Println("nomer ke", nomor)

	}

	//for juga kadang digunakan untuk mengambil data array atau slice
	slice := [...]string{"khoirul", "fikri", "omte"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for kadang digunakan untuk map

	person := map[string]string{
		"name":    "ranse",
		"address": "jl. garuda",
	}
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
