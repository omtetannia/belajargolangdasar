package main

import "fmt"

type filter func(string) string

func sayHelloWithFilter(name string, filter filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)

}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("eko", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)

}
