package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println(total)

	//membuat data slice dan kemudian dijumlahkan dengan menambahkan variabel argumen ... (titik tiga)

	slice := []int{10, 10, 10, 10, 10}
	total = sumAll(slice...)
	fmt.Println(total)

}
