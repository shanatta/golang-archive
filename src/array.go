//Tipe Data Array

package main

import "fmt"

func main() {

	var name [3]string

	name[0] = "Kanaya"
	name[1] = "Azril"
	name[2] = "Baihaqy"

	fmt.Println("==============Array==============")
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)

	fmt.Println("Panjang array name:", len(name))
	fmt.Println("Panjang array vaues:", len(values))

	var Lagi [10]string

	fmt.Println("Panjang array Lagi:", len(Lagi))
}
