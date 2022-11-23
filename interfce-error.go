/**
ERROR INTERFACE
~ Go-Lang memiliki intarface yang digunakan sebagai kontrak untuk membuat error,
nama intarfacenya adalah error
Membuat error
~ Untuk membuat error, kita tidak perlu manual
~ Go-Lang sudah menyediakan libary untuk membuat helper secara mudah, yang terdapat
di package errors
*/

package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian tidak bisa 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}

}

func main() {
	hasil, err := Pembagi(100, 20)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())

	}

}
