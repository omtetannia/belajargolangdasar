/**
POINTER DI METHOD
~ Walaupun method akan menempel di struc, tapi sebenernya data struct yang diakses di dalam method adalah pass by value
~ sangan direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu
diduplikasi ketika memanggil method
*/

package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	data := Man{"omte"}
	data.Married()
	fmt.Println(data)

}
