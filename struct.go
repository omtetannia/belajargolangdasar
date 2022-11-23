package main

import "fmt"

/**
~ struct adalah sebuah tamplate yang digunakan untk menggabungkan data nol atau lebih tipe data dalam satu kesatuan
~ sederhanya struct adalah kumpulan dari field
*/

type Customer struct {
	Name, Address string
	Age           int
}

//struct methode
func (customer Customer) sayHi() {
	fmt.Println("My Name is:", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("form is:", a.Address)
}

func (b Customer) sayAge() {
	fmt.Println("Age is:", b.Age)
}

/**
MEMBUAT DATA DI STRUCT
~ struct juga disebut prototype data, sehingga struct tidak bisa langsung digunakan
~ Namun kita bisa membuat data/object dari struct yang telah kita buat
*/

func main() {
	var data Customer
	data.Name = "Omte"
	data.Address = "Jl. Garuda"
	data.Age = 20

	data.sayHi()
	data.sayHuuu()
	data.sayAge()

	//cara lebih singkat mambuat data di struct
	datadua := Customer{
		Name:    "Joni",
		Address: "Jl. jalan aja",
		Age:     25,
	}

	fmt.Println(data)
	fmt.Println(datadua)

}
