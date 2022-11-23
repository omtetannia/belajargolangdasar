package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings.Contains(string, search) Mengecek apakah string mengandung string lain
	fmt.Println(strings.Contains("omte tania", "omte"))
	fmt.Println(strings.Contains("omte tania", "fikri"))

	//strings.Split(string, separator)  Memotong string berdasarkan separator
	fmt.Println(strings.Split("omte tannia", " "))

	//strings.ToLower(string) Membuat semua karakter string menjadi lower case atau huruf kecil
	fmt.Println(strings.ToLower("Omte Tannia"))
	//strings.ToUpper(string) Membuat semua karakter string menjadi upper case atau huruf besar
	fmt.Println(strings.ToUpper("Omte Tannia"))

	//strings.Trim(string, cutset) Memotong cutset di awal dan akhir string atau menghilangkan spasi di kiri dan kanan
	fmt.Println(strings.Trim("    omte tannia   ", " "))
	fmt.Println(strings.Trim("a    omte tannia   a", " "))

	//strings.ReplaceAll(string, from, to) mengubah semua string dari from ke to
	fmt.Println(strings.ReplaceAll("omte omte omte", "omte", "tannia"))
	fmt.Println(strings.ReplaceAll("omte fikri omte", "omte", "tannia"))

}
