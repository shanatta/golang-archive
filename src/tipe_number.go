//Tipe Number

//2 jenis tipe number (int, floating point/desimal)

//-int
//int8= min(-128), max(127)
//int16= min(-32768), max(32767)
//int32= min(-2147483648), max(2147483647)
//int64= min(-9223372036854775808), max(9223372036854775807)
//uint8= min(0), max(255) (unsigned integer)
//uint16= min(0), max(65535)
//uint32= min(0), max(4294967295)
//uint64= min(0), max(18446744073709551615)

//-floating point
//float32 min(1.18x10^-38), max(3.4x10^38)
//float64 min(2.23x10^-308), max(1.80x10^308)
//complex64
//complex128

//-alias
//byte= uint8
//rune= int32
//int= minim int32 (sesuai os, kalau 64 bit brti otomatis int64)
//uint= minim uint32

//golang case sensitive (syntax harus sama, besar kecilnya)

package main

import "fmt"

func main() {
	fmt.Println("=============Tipe Number=============")
	fmt.Println("Satu= ", 1)
	fmt.Println("Dua= ", 2)
	fmt.Println("Tiga koma lima= ", 3.5)
}
