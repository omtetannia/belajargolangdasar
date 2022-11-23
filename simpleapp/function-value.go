package main

import "fmt"

func getgoodBye(name string) string {
	return "goodBye " + name
}

func main() {

	//func yang dibuat tadi juga bisa di jadikan value
	goodbye := getgoodBye
	result := goodbye("Tannia")
	fmt.Println(result)

}
