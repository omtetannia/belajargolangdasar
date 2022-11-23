/**
Package
~ membuat package caranya adalah membuat folder baru
~ dalam satu package/folder tidak boleh ada function yang sama
~ untuk memanggil package kita bisa menggunakan import
*/

package helper

import "fmt"

func SayHello(Name string) {
	fmt.Println("Hello", Name)
}
