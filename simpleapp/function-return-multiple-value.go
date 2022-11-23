package main

import "fmt"

func getFullName() (string, string, string) {
	return "omte", "tannia", "khoirul"
}

func main() {
	fristname, middlename, lastname := getFullName()
	fmt.Println(fristname, middlename, lastname)

	//jika hanya mengambil bebarapa value maka tinggal di ganti menggunakan _
	//fristname, _, _ := getFullName()
	//fmt.Println(fristname)
}
