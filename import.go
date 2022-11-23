/**
Access Modifier
~ Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
~ Jika nama nya diawali dengan huruf besar maka artinya bisa diakses dara package lain, jika dimulai
dengan huruf kecil artinya tidak bisa diakses dari package lain
~ sederhanya jika di awali huruf besar maka bisa di akses dari luar (public) jika di awali huruf kecil
tidak bisa di akses dari luar (private)
*/

package main

import "belajargolangdasar/helper"

func main() {
	helper.SayHello("Omte")
}
