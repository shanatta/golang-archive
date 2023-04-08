//Variable

//tempat menyimpan data
//di golang harus tipe data sama, bikin beberapa variable
//1 variable unik, kalau a= string, bawahnya a=int gabisa

package main

import "fmt"

func main() {
	var name string

	fmt.Println("===============Variable===============")
	name = "Ainz Ooalgown"
	fmt.Println(name)

	name = "Aizen Sosuke"
	fmt.Println(name)

	///name = true
	///fmt.Println(name) (error karena beda tipe data)

	fmt.Println("--------------------------------------")
	var name2 = "Dazai Osamu" //var doang bisa, tanpa sebutin tipe datanya
	fmt.Println(name2)

	name2 = "Azuru"
	fmt.Println(name2)

	fmt.Println("--------------------------------------")
	var boolean = true //tipe data boolean
	fmt.Println(boolean)

	boolean = false
	fmt.Println(boolean)

	fmt.Println("--------------------------------------")
	var umur = 17 //tipe data int
	fmt.Println(umur)

	umur = 18
	fmt.Println(umur)

	fmt.Println("--------------------------------------")
	angka := 28.12 //tipe data floatpoint //var bisa diganti :=
	fmt.Println(angka)

	angka = 29.12
	fmt.Println(angka)

	var ( //banyak variable
		firstname = "Azril"
		lastname  = "Baihaqy"
	)

	fmt.Println("--------------------------------------")
	fmt.Println(firstname)
	fmt.Println(lastname)
}
