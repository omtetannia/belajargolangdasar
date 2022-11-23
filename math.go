/*
math.Round(float64) Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
math.Floor(float64) Membulatkan kebawah
math.Ceil(float64) Membulatkan float64 keatas
math.Max(float64, float64) mengemnalikan nilai floa paling besar
math.Min(float64, float64) mengembalikan nilai float paling kecil
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))
	fmt.Println(math.Floor(2.5))
	fmt.Println(math.Ceil(2.5))
	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))

}
