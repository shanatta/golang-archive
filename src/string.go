//Tipe Data String

package main

import "fmt"

func main() {

	fmt.Println("================String================")
	fmt.Println("Okaeri")
	fmt.Println("Doktah")
	fmt.Println("Lopyu")

	fmt.Println("=============Index String=============")
	fmt.Println(len("Okaeri")) //jadi 6 karena dimulai dari 1, spasi&tanda lain dihitung
	fmt.Println("Doktah"[0])   //jadi 68= byte nya d
	fmt.Println("Lopyu"[4])    //jadi 117= byte nya u
}
