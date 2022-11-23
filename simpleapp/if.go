package main

import "fmt"

func main() {
	var name = "omte"

	if name == "omte" {
		fmt.Println("Hello omte")
	} else {
		fmt.Println("maaf, anda siapa")
	}
	//else jika ingin eksekusi bernialai false

	name = "ani"

	//menggabungkan if dan else

	if name == "ani" {
		fmt.Println("Hello ani")
	} else if name == "Rusid" {
		fmt.Println("Hello Rusid")
	} else {
		fmt.Println("maaf")
	}

	//membuat statement sederhana sebelum melakukan pengecekan terhadap kondisi

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama valid")
	}

}
