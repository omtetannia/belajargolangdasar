package main

import "fmt"

func main() {

	var name = "tannia"

	switch name {
	case "omte":
		fmt.Println("Hello omte")
	case "tannia":
		fmt.Println("Hello tannia")
	default:
		fmt.Println("silahkan registrasi")
	}

	// short statement berfungsi menambahakan statement sebelum swicth

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("valid")
	case false:
		fmt.Println("Tidak Valid")
	}
}
