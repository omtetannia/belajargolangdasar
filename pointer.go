/*
*
POINTER
Pass by Value
~ Secara default Go-Lang semua variable itu di passing by value, bukan by reference
~ Artinya, Jika kita mengirim sebuah variabel ke dalam function, method atau variable lain, sebenernya yang dikirim
adalah duplikasi valuenya
*/
package main

import "fmt"

type Address struct {
	City, Provinsi, Country string
}

func main() {
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	//Operator & (operator and) untuk pointer pada variable pada value field
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"

	//operator * (operator bintang) untuk pointer pada variable semua value field
	*address3 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	//Function New membuat pointer dengan data kosong
	address4 := new(Address)
	//address4.City = "jakarta" //jika diisi
	fmt.Println(address4)

}

/**
POINTER DI FUNCTION
~ Saat kita membuat parameter di functio, secara default adalah pass by value,
artinya data akan di copy lalu dikirim ke function tersebut
~ Oleh karena itu, jika kita mengubah data dalam function, data yang aslinya tidak akan berubah
~ Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
~ Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
~ Untuk melakukan ini kita juga bisa menggunakan ponter di function
~ Untuk membuat sebuah parameter sebagai pointer, Kita bisa menggunakan Operator* di parameternya
*/
