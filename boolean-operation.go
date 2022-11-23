package main

import "fmt"

func main() {
	var (
		Nilaiakhir      = 80
		NilaiAbsen      = 80
		LulusNilaiAkhir = Nilaiakhir >= 80
		LuluNilaiAbsen  = NilaiAbsen >= 80
	)

	var Lulus = LulusNilaiAkhir && LuluNilaiAbsen

	fmt.Println(LulusNilaiAkhir)
	fmt.Println(LuluNilaiAbsen)
	fmt.Println(Lulus)

	//cara cepat
	var (
		nilaiawal  = 70
		nilailulus = 80
	)

	fmt.Println(nilaiawal >= 80 && nilailulus >= 80)
}
