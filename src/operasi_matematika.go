//Operasi Matematika

package main

import "fmt"

func main() {

	var result = 28 + 29
	fmt.Println("========Operasi Matematika========")
	fmt.Println("Hasil 28+29=", result)

	var a = 28
	var b = 29
	var c = a * b
	fmt.Println("Hasil 28*29=", c)

	var i = 2
	i += 12 //i= i + 12

	fmt.Println("======Augmented Assignments======")
	fmt.Println("i= 2")
	fmt.Println("Hasil i+12=", i)

	i++ //i= i + 1
	fmt.Println("==========Unary Operator==========")
	fmt.Println("i= 14")
	fmt.Println("Hasil i+1=", i)

	i-- //i= i - 1
	fmt.Println("\ni= 15")
	fmt.Println("Hasil i-1=", i)

	var negative = -2006
	var positive = +2005
	fmt.Println("\nNilai negatif:", negative)
	fmt.Println("Nilai positif:", positive)
}

/*Augmented Assignments
a= a+10
a += 10

a= a-10
a -= 10

a= a*10
a *= 10

a= a/10
a /= 10

a= a%10
a %= 10

-Unary Operator
++ (a= a+1)
-- (a= a-1)
- (Negative)
+ (Positive)
! (Boolean Kebalikan)*/
