package main

import "fmt"

func main() {

	var DataString string
	DataString = "omte"
	fmt.Println(DataString)

	DataString = "Tannia"
	fmt.Println(DataString)

	//type data pada var tidak harus mendeklarasikan type data
	var name = "khoirul"
	fmt.Println(name)

	name = "fikri"
	fmt.Println(name)

	var age = 17
	fmt.Println(age)

	age = 18
	fmt.Println(age)

	//kata kunci var bisa di rubah dengan := di awal, kata kunci := hanya di gunakan di awal untuk setelhanya bisa menggunakan =
	country := "indonesia"
	fmt.Println(country)

	country = "thailand"
	fmt.Println(country)

	country = "singapore"
	fmt.Println(country)

	country = "cina"
	fmt.Println(country)

	country = "india"
	fmt.Println(country)

	//multi variable
	var (
		firstname = "rois"
		lastname  = "alam"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

}
