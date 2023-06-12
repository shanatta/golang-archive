//Konversi Tipe Data

// -konversi tipe data 1
package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("============Konversi 1============")
	fmt.Println("Nilai int32=", nilai32)
	fmt.Println("Nilai int64=", nilai64)
	fmt.Println("Nilai int8=", nilai8)

	var name = "Azuru"
	var a byte = name[0]
	var z byte = name[1]
	var aString string = string(a)
	var zString string = string(z)

	fmt.Println("============Konversi 2============")
	fmt.Println(name)
	fmt.Println(aString)
	fmt.Println(zString)
}
