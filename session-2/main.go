package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// * Temporary variable
	// * if else
	// var currentYear = 2021

	// if age := currentYear - 1998; age < 17{
	// 	fmt.Println("Kamu belum boleh membuat kartu sim")
	// 	} else {
	// 	fmt.Println("Kamu sudah boleh membuat kartu sim")
	// }

	// * switch
	// var score = 8

	// switch score {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }
	
	// var score = 4
	
	// switch {
	// case score == 8:
	// 	fmt.Println("perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not bad")
	// 	fallthrough
	// case score < 5:
	// 	fmt.Println("It is ok, but please study harder")
	// default:
	// 	{
	// 		fmt.Println("study harder")
	// 		fmt.Println("You don't have a good score yet")
	// 	}
	// }

	// * nested conditions
	// var score = 10

	// if score > 7 {
	// 	switch score {
	// 	case 10:
	// 		fmt.Println("perfect!")
	// 	default:
	// 		fmt.Println("nice")
	// 	}
	// } else {
	// 	if score == 5 {
	// 		fmt.Println("not bad")
	// 	} else if score == 3 {
	// 		fmt.Println("keep trying")
	// 	} else {
	// 		fmt.Println("you can do it")
	// 		if score == 0 {
	// 			fmt.Println("try harder!")
	// 		}
	// 	}
	// }

	// * for looping
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// * while looping
	// var i = 0

	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// * break looping
	// var i = 0

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++

	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// * break and continue keyword
	// for i := 1; i <= 10; i++ {
	// 	if i % 2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// * nested looping
	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println() // == fmt.Print("\n") == 
	// }

	// * looping label
	// outerLoop:
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Perulangan ke - n", i + 1)
	// 	for j := 0; j < 3; j++ {
	// 		if i == 2 {
	// 			break outerLoop
	// 		}
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Print("\n")
	// }

	// * array
	// var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"Airell", "Nanda", "Mailo"}

	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// * modify array through index
	// var fruits = [3]string{"apel", "pisang", "mangga"}
	// fruits[0] = "apple"
	// fruits[1] = "pisang"
	// fruits[2] = "mangga"

	// fmt.Printf("%#v\n", fruits)

	// * array looping
	// var fruits = [3]string{"apple", "banana", "mango"}

	// // Cara Pertama
	// for i, v := range fruits {
	// 	fmt.Printf("Index: %d, Value %s\n", i, v)
	// }

	// fmt.Println(strings.Repeat("#", 25))

	// // Cara kedua
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	// }

  // * multidimensional array
	// balances := [2][3]int{{5,6,7},{8,9,10}}

	// for _, arr := range balances {
	// 	for _, value := range arr {
	// 		fmt.Printf("%d ", value)
	// 	}
	// 	fmt.Println()
	// }

	// * slice
	// var fruits = []string{"apple", "banana", "mango"}

	// _ = fruits

	// fmt.Printf("%#v", fruits)

	// * create slice using make function
	// var fruits = make([]string, 3)

	// _ = fruits

	// fmt.Printf("%#v", fruits)

	// * add element to slice using append function
  // var fruits = make([]string, 3)
	// _ = fruits

	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	// fruits = append(fruits, "apple", "banana", "mango")

	// fmt.Printf("%#v", fruits)

	// * add element to slice using append function with ellipsis
	// var fruits1 = []string{"apple", "banana", "manggo"}
	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// fruits1 = append(fruits1, fruits2...)

	// fmt.Printf("%#v\n", fruits1)
	// fmt.Printf("%#v\n", fruits2)

	// * copy slice using copy function
	// var fruits1 = []string{"apple", "banana", "mango"}
	// var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// nn := copy(fruits1, fruits2)

	// fmt.Println("Fruits1 =>", fruits1)
	// fmt.Println("Fruits2 =>", fruits2)
	// fmt.Println("Copied elements =>", nn)

	// * slicing
	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	// var fruits2 = fruits1[1:4]
	// fmt.Printf("%#v\n", fruits2)

	// var fruits3 = fruits1[0:]
	// fmt.Printf("%#v\n", fruits3)

	// var fruits4 = fruits1[:3]
	// fmt.Printf("%#v\n", fruits4)

	// var fruits5 = fruits1[:] // sama dengan nilai fruits[:len(fruits1)]
	// fmt.Printf("%#v\n", fruits5)

	// * combining slicing and append
	// var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	// fruits1 = append(fruits1[:3],"rambutan")

	// fmt.Printf("%#v\n", fruits1)

	// * backing array
	// var fruits1 = []string{"apple", "mango", "durian", "banana", "starfruit"}

	// var fruits2 = fruits1[2:4]

	// fruits2[0] = "rambutan"

	// fmt.Println("fruits1 =>", fruits1)
	// fmt.Println("fruits2 =>", fruits2)

	// * cap function
	// var fruits1 = []string{"apple", "mango", "durian", "banana"}

	// fmt.Println("Fruits1 cap:", cap(fruits1)) // 4
	// fmt.Println("Fruits1 len:", len(fruits1)) // 4

	// fmt.Println(strings.Repeat("#", 20))
	
	// var fruits2 = fruits1[0:3]
	// fmt.Println("Fruits2 cap:", cap(fruits2)) // 4
	// fmt.Println("Fruits2 len:", len(fruits2)) // 3

	// fmt.Println(strings.Repeat("#", 20))

	// var fruits3 = fruits1[1:]
	// fmt.Println("Fruits3 cap:", cap(fruits3)) // 3
	// fmt.Println("Fruits3 len:", len(fruits3)) // 3


	// * create a new backing array
	// cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	// newCars := []string{}

	// newCars = append(newCars, cars[0:2]...)

	// cars[0] = "Nissan"
	// fmt.Println("cars:", cars)
	// fmt.Println("newCars:", newCars)

	// * map
	// var person map[string]string // Deklarasi

	// person = map[string]string{} // Inisialisasi

	// person["name"] = "Pratama"

	// person["age"] = "23"

	// person["address"] = "Perumnas Kepanjen"

	// fmt.Println("name:", person["name"])
	// fmt.Println("age:", person["age"])
	// fmt.Println("address:", person["address"])

	// var person = map[string]string{
	// 	"name"		: "Pratama",
	// 	"age"			:	"23",
	// 	"address"	: "Perumnas Kepanjen",
	// }

	// fmt.Println("name:", person["name"])
	// fmt.Println("age:", person["age"])
	// fmt.Println("address:", person["address"])

	// * map looping
	// var person = map[string]string{
	// 	"name"		: "Pratama",
	// 	"age"			:	"23",
	// 	"address"	: "Perumnas Kepanjen",
	// }

	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	// * delete value in map
	// var person = map[string]string{
	// 	"name"		: "Pratama",
	// 	"age"			:	"23",
	// 	"address"	: "Perumnas Kepanjen",
	// }

	// fmt.Println("Before deleting:", person)

	// delete(person, "age")

	// fmt.Println("After deleting:", person)

	// * detect value in map
	// var person = map[string]string{
	// 	"name"		: "Pratama",
	// 	"age"			:	"23",
	// 	"address"	: "Perumnas Kepanjen",
	// }

	// value, exist := person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value doesn't exist")
	// }

	// delete(person, "age")

	// value, exist = person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value has been deleted")
	// }

	// * combining slice with map
	// var people = []map[string]string{
	// 	{"name": "Airell", "age": "23"},
	// 	{"name": "Nanda", "age": "23"},
	// 	{"name": "Mailo", "age": "15"},
	// }

	// for i, person := range people {
	// 	fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	// }

	// * aliase
	// // deklarasi variable dengan tiped ata uint8
	// var a uint8 = 10
	// var b byte // byte adalah alias dari tipe data uint8

	// b = a
	// _ = b

	// // Mendeklarasi tipe data alias bernama second
	// // type nama_alias = nama_tipe_data
	// type second = uint
	// var hour second = 3600
	// fmt.Printf("hour type: %T\n", hour) // => hour type : uint

	// * string in depth
	// * looping over string (byte-to-byte)
	// city := "Jakarta"

	// for i := 0; i < len(city); i++ {
	// 	fmt.Printf("index: %d, byte: %d\n", i, city[i])
	// }

	// var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	// fmt.Println(string(city))
	// fmt.Println(city)

	// * looping over string (rune-by-rune)
	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte length ==> %d\n", len(str1)) // str1 byte length ==> 5
	fmt.Printf("str2 byte length ==> %d\n", len(str2)) // str2 byte length ==> 6

	fmt.Printf("str1 character length ==> %d\n", utf8.RuneCountInString(str1)) // str1 character length ==> 5
	fmt.Printf("str2 character length ==> %d\n", utf8.RuneCountInString(str2)) // str2 character length ==> 5

	str := "mânca"

	for i, s := range str {
		fmt.Printf("index ==> %d, string %s\n", i, string(s))
	}

	/*
	index ==> 0, string m
	index ==> 1, string â
	index ==> 3, string n
	index ==> 4, string c
	index ==> 5, string a
	*/
}
