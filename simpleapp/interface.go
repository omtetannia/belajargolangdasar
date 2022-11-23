/**
INTERFACE
~ Interface adakah tyoe data yang abtrack,  dia tidak memiliki implementasi langsung
~ sebuah intreface berisikan definisi-definsi method/fuction
~ Biasanya Interface digunakan sebagai kontrak
IMPLEMENTASI INTERFACE
~ Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai intarface tersebut
~ Sehingga kita tidak perlu mengimplementasikan interface secara manual
~ Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat intarface, kita harus menyebutkan secara eksplisit
akan menggunakan interface mana
*/

package main

import "fmt"

type HasName interface {
	GetName() string
}

func Hello(hasname HasName) {
	fmt.Println("Helloo", hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

//contoh 2
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var Omte Person
	Omte.Name = "Omte"

	Hello(Omte)

	cat := Animal{
		Name: "Kucing",
	}
	Hello(cat)
}
