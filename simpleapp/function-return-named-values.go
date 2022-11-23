package main

import "fmt"

func getFullName() (NamaDepan string, NamaBelakang string) {
	NamaDepan = "Khoirul"
	NamaBelakang = "Fikri"

	return
}

func main() {
	firstName, Lastname := getFullName()
	fmt.Println(firstName)
	fmt.Println(Lastname)

}
