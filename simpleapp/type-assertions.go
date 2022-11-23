/**
Type Assertions
~ Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
~ Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
Type Assertions menggunakan switch
~ Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic diaplikasi kita
~ Jika panic dan tidak ter-cover, maka otomatis program kita akan mati
~ Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
*/

package main

import "fmt"

func Random() interface{} {
	return 100
}

func main() {
	result := Random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	//contoh Type Assertions menggunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("is string", value)
	case int:
		fmt.Println("is int", value)
	default:
		fmt.Println("uknown")
	}

}
