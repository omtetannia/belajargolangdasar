package main

import (
	"fmt"
)

// recover untuk aplikasi tetap berjalan meski respon panic/error, recover di taruh di dalam defer function
func endApp() {
	fmt.Println("End App")
	message := recover()
	if message == nil {
		fmt.Println("response meesage", message)
	}

}

func runApp(error bool) {
	//ada system error dan mau meberhentikannnya maka menggunakan panic
	defer endApp()
	if error {
		panic("error")
	}

	fmt.Println("aplikasi berjalan")

}

func main() {
	runApp(true)

}
