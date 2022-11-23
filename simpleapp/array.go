package main

import "fmt"

func main() {

	var names [4]string
	names[0] = "omte"
	names[1] = "tannia"
	names[2] = "fikri"
	names[3] = "khoirul"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	//cara cepat
	//jika mengetahui datanya ada 3 maka bisa menggunakan [3] namun jika tidak tau ada berapa bisa menggunakan [...]
	var city = [3]string{
		"jogja",
		"lumajang",
		"malang",
	}

	fmt.Println(city)
	fmt.Println(city[0])
	fmt.Println(city[1])
	fmt.Println(city[2])
}
