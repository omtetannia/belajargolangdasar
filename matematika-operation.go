package main

import "fmt"

func main() {
	var (
		a = 10
		b = 20
		c = a + b
	)

	// cara langsung emggunaka result tanpa memisah variable nay
	var perkurangan = 10 - 5
	var perkalian = 10 * 5

	//augmented assignments
	var hasilperkalian = 50
	hasilperkalian += 10

	//unary operation
	var j = 80
	j++ //artinga 1 + 80

	fmt.Println(perkurangan)
	fmt.Println(perkalian)
	fmt.Println(hasilperkalian)
	fmt.Println(j)
	fmt.Println(c)

}
