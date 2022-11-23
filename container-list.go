package main

/*
Package contaner/list
~ Package container/list adalah implementasi struktu data double linked list di Go-Lang
*/

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("omte")
	data.PushBack("tannia")
	data.PushBack("fikri")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value)
	}
}
