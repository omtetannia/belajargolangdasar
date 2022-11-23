package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args

	fmt.Println("Argumen : ", arg)

	Name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", Name)

	} else {
		fmt.Println("Erorr", err.Error())
	}
}
