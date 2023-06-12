//Type Deklarasi

//-kemampuan membuat ulang tipe data baru dari tipe yg sudah ada
//-biasanya digunakan untuk membuat alias agar lebih mudah dimengerti

package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAzu NoKTP = "2812200529"
	var marriedStatus Married = false
	fmt.Println("========Type Declarations========")
	fmt.Println("No KTP Azu:", noKtpAzu)
	fmt.Println("Married Status:", marriedStatus)
}
