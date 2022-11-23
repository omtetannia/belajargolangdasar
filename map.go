package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Omte Tannia",
		"address": "jl. garuda",
	}

	//menambah map pada person

	person["status"] = "kawin"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Print(person["address"])

	fmt.Println(person["status"])

}
