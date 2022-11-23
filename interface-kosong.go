/*
*
INTERFACE KOSONG
~ Golang bukanlah bahasa pemrograman yang berorientasi objek
~ Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa
dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
~ Contoh di Java ada java.lang.Object
~Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
~ Interface kosong adalah interface yang tidak memiliki deklarasai method satupun,
hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
~ contoh penggunaan intarface kosong di Go-Lang

	fmt.Println(a ... interface{})
	panic(v interface{})
	recover() interface{}
	dan lain-lain
*/
package main

import "fmt"

func Ups() interface{} {
	return "Ups"
}

//menggunakan parameter saat membuat interface kosong
func ParameterUps(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Parameter Ups"
	}
}

func main() {
	kosong := Ups()
	parameter := ParameterUps(5)
	fmt.Println(parameter)
	fmt.Println(kosong)
}
