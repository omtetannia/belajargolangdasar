package main

import "fmt"

func main() {

	const firstName string = "khoirul"
	const lastName = "fikri"

	fmt.Println(firstName)
	fmt.Println(lastName)

	//multiple constant
	const (
		FirstName = "aris"
		LastName  = "aji"
	)

	fmt.Println(FirstName)
	fmt.Println(LastName)

}
