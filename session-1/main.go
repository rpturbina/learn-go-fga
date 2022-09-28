package main

import "fmt"

func main() {
	// fmt.Println("Hello World", "Nice", "Wow")
	// fmt.Print("Hello World", "Nice", "Wow")

	// * Variable declaration with data type
	// var name string = "Pratama"
	// var age int = 23

	// name = "Rizki"
	// age = 25

	// fmt.Println("Ini adalah namanya ===>", name)
	// fmt.Println("Ini adalah umurnya ===>", age)

	// * Variable decalaration without data type
	// var name = "Pratama"
	// var age = 23

	// * Short declaration
	// name := "Pratama"
	// age := 23

	// fmt.Printf("%T, %T", name, age)

	// * Multiple variable declarations
	// var student1, student2, student3 string = "satu", "dua", "tiga"
	
	// var first, second, third int

	// first, second, third = 1, 2, 3

	// fmt.Println(student1, student2, student3)

	// fmt.Println(first, second, third)

	// var name, age, address = "Pratama",  23, "Perumnas"

	// first, second, third := "1", 2, "3"

	// fmt.Println(name, age, address)

	// fmt.Println(first, second, third)

	// * Underscore variable
	// var firstVariable string

	// var name, age, address = "Pratama", 23, "Perumnas Kepanjen"

	// _, _, _, _ = firstVariable, name, age, address

	// * fmt.Printf Usage
	
	// first, second := 1, "2"

	// fmt.Printf("Tipe data variabel first adalah %T \n", first)
	// fmt.Printf("Tipe data variabel second adalah %T \n", second)

	// var name = "Pratama"

	// var age = 23

	// var address = "Perumnas Kepanjen"

	// fmt.Printf("Halo nama ku %s, umur aku adalah %d, dan aku tinggal di %s \n", name, age, address)

	// * data type : Number (integer)
	// var first = 89
	// var second = -12

	// var first uint8 = 89
	// var second int8 = -12

	// fmt.Printf("tipe data first: %T\n", first)
	// fmt.Printf("bilangan second: %T\n", second)

	// * data type : Number (float)
	// var decimalNumber float32 = 3.63

	// fmt.Printf("decimal number: %f\n", decimalNumber)
	// fmt.Printf("decimal number: %.3f\n", decimalNumber)

	// * data type : Bool
	// var condition bool = true
	// fmt.Printf("is it permitted? %t \n", condition)

	// * data type : String
	// var message = "String"
	// fmt.Print(message)

// 	var message = `Selamat datang di "Hacktiv8".
// Salam kenal :).
// Mari belajar "Scalable Web Service With Go".
// `

// 	fmt.Print(message)

	// * data type: nil & zero value


	// * constant

	// const full_name string = "Rizki Pratama Turbina"

	// fmt.Printf("Hello %s", full_name)

	// * Operators (Arithmetic Operators)
	// var value = 2 + 2 * 3
	// fmt.Println(value)

	// var value_2 = (2 + 2) * 3
	// fmt.Println(value_2)

	// * Operators (Relational Operators)
	// var firstCondition bool = 2 < 3
	// var secondCondition bool = "joey" == "Joey"
	// var thirdCondition bool = 10 != 2.3
	// var fourthCondition bool = 11 <= 11

	// fmt.Println("first condition:", firstCondition)
	// fmt.Println("second condition:", secondCondition)
	// fmt.Println("third condition:", thirdCondition)
	// fmt.Println("fourth condition:", fourthCondition)

	// * Logical Operators
	var right = true
	var wrong = false
	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}