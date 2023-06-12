//Operasi Boolean

/*
Operasi &&
true && true = true
true && false = false
false && true = false
false && false = false

Operasi ||
true || true = true
true || false = true
false || true = true
false || false = false

Operasi !
!true = false
!false = true
*/
package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println("=========Operasi Boolean=========")
	fmt.Println("Lulus Ujian:", lulusUjian)
	fmt.Println("Lulus Absensi:", lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println("Hasil:", lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)
}
