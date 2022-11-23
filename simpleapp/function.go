package main

import "fmt"

func sayHello() {
	fmt.Println("Hello Word")
}

func main() {
	sayHello()

	//bisa dengan perulangan
	for i := 0; i < 10; i++ {
		sayHello()
	}

}
