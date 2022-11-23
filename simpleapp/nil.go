/*
*
NIL
~ Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi
maka secara otomatis nilainya adalah null atau nil
~ Berbeda dengan Go-Lang, di Go-Lang Saat kita buat variable dengan tipe data tertentu,
maka secara otomatis akan dibuatkan default value nya
~ Namun di Go-Lang ada data nil, Yaitu data kosong
~ Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function,
map, slice, pointer dan channel
*/
package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("omte")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}

}
