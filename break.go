package main

import "fmt"

func main() {
	for index := 0; index < 10; index++ {
		if index == 5 {
			break
		}
		fmt.Println("perulangan ke", index)
	}
}

//digunakan untuk memberhentikan perulangan
