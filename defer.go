package main

import "fmt"

func Logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication(value int) {

	//fungsi defer adalah tetap memanggil function yang di inginkan meskipun terjadi error
	defer Logging()
	fmt.Println("run application")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApplication(10)

}
