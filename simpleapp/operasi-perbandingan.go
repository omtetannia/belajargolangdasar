package main

import "fmt"

func main() {
	var (
		name1            = "omte"
		name2            = "tannia"
		resultsamadengan = name1 == name2
	)

	//Lebih dari >
	var (
		name3           = "omte"
		name4           = "tannia"
		resultlebihdari = name3 > name4
	)
	//Kurang dari <
	var (
		name5            = "omte"
		name6            = "tannia"
		resultkurangdari = name5 < name6
	)
	//Lebih dari sama dengan >=
	var (
		name7                     = "omte"
		name8                     = "tannia"
		resultlebihdarisamadengan = name7 >= name8
	)
	//Kurang dari sama dengan <=
	var (
		name9                      = "omte"
		name10                     = "tannia"
		resultkurangdarisamadengan = name9 < name10
	)

	fmt.Println(resultsamadengan)
	fmt.Println(resultlebihdari)
	fmt.Println(resultkurangdari)
	fmt.Println(resultlebihdarisamadengan)
	fmt.Println(resultkurangdarisamadengan)

	//menggunakan angka

	var (
		data1 = 20
		data2 = 10
	)

	fmt.Println(data1 > data2)
	fmt.Println(data1 < data2)
	fmt.Println(data1 >= data2)
	fmt.Println(data1 <= data2)
}
