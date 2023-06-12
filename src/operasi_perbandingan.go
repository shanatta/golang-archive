//Operasi Perbandingan

package main

import "fmt"

func main() {

	var name1 = "Azuru"
	var name2 = "azuru"

	var result bool = name1 == name2
	fmt.Println("=======Operasi Perbandingan=======")
	fmt.Println("Data 1:", result)

	var value1 = 28
	var value2 = 29

	fmt.Println("Data 2:", value1 > value2)
	fmt.Println("Data 3:", value1 < value2)
	fmt.Println("Data 4:", value1 >= value2)
	fmt.Println("Data 5:", value1 <= value2)
	fmt.Println("Data 6:", value1 == value2)
	fmt.Println("Data 7:", value1 != value2)
}
