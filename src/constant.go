//Constant

// -variabel tidak bisa diubah
package main

import "fmt"

/*-constant 1 1
func main() {
	const firstName string = "Azril"
	const lastName = "Baihaqy"
	const value = 28

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	//-error
	firstName = "Tidak Bisa Diubah"
	lastName = "Tidak Bisa Diubah"
}*/

//-constant langsung banyak
func main() {
	const (
		firstName string = "Nat"
		lastName         = "Not"
		value            = 29
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
