package main

import "fmt"

// * function
// func main() {
// 	greet("Pratama", 23)
// }

// func greet(name string, age int8) {
// 	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
// }

// func main() {
// 	greet("Pratama", "Perumnas Kepanjen")
// }

// func greet(name, address string) {
// 	fmt.Println("Hello there! My name is", name)
// 	fmt.Println("I live in", address)
// }

// * function return
// func main() {
// 	var names = []string{"Airell", "Jordan"}
// 	var printMsg = greet("Heii", names)

// 	fmt.Println(printMsg)
// }

// func greet(msg string, names []string) string {
// 	var joinStr = strings.Join(names, " ")

// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)

// 	return result
// }

// * function returning multiple values
// func main() {
// 	var diameter float64 = 15

// 	var area, circumference = calculate(diameter)

// 	fmt.Println("Area:", area)
// 	fmt.Println("Circumference:", circumference)
// }

// func calculate(d float64) (float64, float64) {
// 	// Menghitung luas
// 	var area float64 = math.Pi * math.Pow(d/2, 2)

// 	// Menghitung keliling
// 	var circumference = math.Pi * d

// 	return area, circumference
// }

// * function predefined return value
// func main() {
// 	var diameter float64 = 15

// 	var area, circumference = calculate(diameter)

// 	fmt.Println("Area:", area)
// 	fmt.Println("Circumference:", circumference)
// }

// func calculate(d float64) (area float64, circumference float64) {
// 	// Menghitung luas
// 	area = math.Pi * math.Pow(d/2, 2)

// 	// Menghitung keliling
// 	circumference = math.Pi * d

// 	return
// }

// * variadic function #1
// func main() {
// 	studentList := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")

// 	fmt.Printf("%v", studentList)
// }

// func print(names ...string) []map[string]string {
// 	var result []map[string]string

// 	for i, v := range names {
// 		key := fmt.Sprintf("student%d", i+1)
// 		temp := map[string]string {
// 			key: v,
// 		}
// 		result = append(result, temp)
// 	}

// 	return result
// }

// * variadic function #2
// func main() {
// 	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

// 	result := sum(numberLists...)
// 	fmt.Println("Result:", result)
// }

// func sum(numbers ...int) int {
// 	total := 0

// 	for _, v := range numbers {
// 		total += v
// 	}

// 	return total
// }

// * variadic function 3
// func main() {
// 	profil("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")
// }

// func profil(name string, favFoods ...string) {
// 	mergeFavFoods := strings.Join(favFoods, ", ")

// 	fmt.Println("Hello there!!! I'm", name)
// 	fmt.Println("I really love to eat", mergeFavFoods)
// }

// * closure : declare closure in variable
// func main() {
// 	var evenNumbers = func(numbers ...int) []int {
// 		var result []int

// 		for _, v := range numbers {
// 			if v % 2 == 0 {
// 				result = append(result, v)
// 			}
// 		}

// 		return result
// 	}

// 	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

// 	fmt.Println(evenNumbers(numbers...))
// }

// * closure : IIFE(immediately-invoked function expression)
// func main() {
// 	var isPalindrome = func(str string) bool {
// 		var temp string = ""

// 		for i := len(str) - 1; i >= 0; i-- {
// 			temp += string(byte(str[i]))
// 		}

// 		return temp == str
// 	}("katak")

// 	fmt.Println(isPalindrome)
// }

// * closure : closure as return value
// func main() {
// 	var studentLists = []string{"Airell", "Nanda", "Mailo", "Schannel", "Marco"}

// 	var find = findStudent(studentLists)

// 	fmt.Println(find("airell"))
// }

// func findStudent(students []string) func(string) string {

// 	return func(s string) string {
// 		var student string
// 		var position int

// 		for i, v := range students {
// 			if strings.ToLower(v) == strings.ToLower(s) {
// 				student = v
// 				position = i
// 				break
// 			}
// 		}

// 		if student == "" {
// 			return fmt.Sprintf("%s doesn't exist!!!", s)
// 		}
// 		return fmt.Sprintf("We found %s at the position %d", s, position + 1)
// 	}
// }

// * closure : callback
// type isOddNum func(int) bool

// func main() {
// 	var numbers = []int{2, 5, 8, 10, 3, 99, 23}

// 	var find = findOddNumbers(numbers, func(number int) bool {
// 		return number % 2 != 0
// 	})

// 	fmt.Println("Total odd numbers", find)
// }

// func findOddNumbers(numbers []int, callback isOddNum) int {
// 	var totalOddNumbers int
// 	for _, v := range numbers {
// 		if callback(v) {
// 			totalOddNumbers++
// 		}
// 	}
// 	return totalOddNumbers
// }

// * pointer
// func main() {
// 	var firstNumber *int
// 	var secondNumber *int
// 	_, _ = firstNumber, secondNumber
// }

// * pointer (memory address)
// func main() {
// 	var firstNumber int = 4
// 	var secondNumber *int = &firstNumber

// 	fmt.Println("firstNumber (value)		:", firstNumber)
// 	fmt.Println("firstNumber (memory address)	:", &firstNumber)

// 	fmt.Println("secondNumber (value)		:", *secondNumber)
// 	fmt.Println("secondNumber (memory address)	:", secondNumber)
// }

// * changing value through pointer
// func main() {
// 	var firstPerson string = "John"
// 	var secondPerson *string = &firstPerson

// 	fmt.Println("firstPerson (value)	:", firstPerson)
// 	fmt.Println("firstPerson (memory address)	:", &firstPerson)

// 	fmt.Println("secondPerson (value)	:", *secondPerson)
// 	fmt.Println("secondPerson (memory address)	:", secondPerson)

// 	fmt.Println("\n",strings.Repeat("#", 30), "\n")

// /	*secondPerson = "Doe"

// 	fmt.Println("firstPerson (value)	:", firstPerson)
// 	fmt.Println("firstPerson (memory address)	:", &firstPerson)

// 	fmt.Println("secondPerson (value)	:", *secondPerson)
// 	fmt.Println("secondPerson (memory address)	:", secondPerson)
// }

// * pointer as a parameter
// func main() {
// 	var a int = 10

// 	fmt.Println("Before:", a)

// 	changeValue(&a)

// 	fmt.Println("After:", a)
// }

// func changeValue(number *int) {
// /	*number = 20
// }

// * struct
// type Employee struct {
// 	name			string
// 	age				int
// 	division	string
// }

// func main() {

// }

// * struct : giving value to struct
// type Employee struct {
// 	name 			string
// 	age				int
// 	division	string
// }

// func main() {
// 	var employee Employee

// 	employee.name = "Pratama"
// 	employee.age = 23
// 	employee.division = "Frontend Engineer"

// 	fmt.Println(employee.name)
// 	fmt.Println(employee.age)
// 	fmt.Println(employee.division)
// }

// * struct : initializing struct
// type Employee struct {
// 	name			string
// 	age				int
// 	division	string
// }

// func main() {
// 	var employee = Employee{}

// 	employee.name = "Pratama"
// 	employee.age = 23
// 	employee.division = "Frontend Engineering"

// 	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

// 	fmt.Printf("Employee1: %+v\n", employee)
// 	fmt.Printf("Employee2: %+v\n", employee2)
// }

// * struct : pointer to a struct
// type Employee struct {
// 	name			string
// 	age				int
// 	division	string
// }

// func main() {
// 	var employee = Employee{name: "Pratama", age: 23, division: "Frontend Engineering"}

// 	var employee2 *Employee = &employee

// 	fmt.Println("Employee1 name:", employee.name)
// 	fmt.Println("Employee2 name:", employee2.name)

// 	fmt.Println(strings.Repeat("#", 20))

// 	employee2.name = "Fadhila"

// 	fmt.Println("Employee1 name:", employee.name)
// 	fmt.Println("Employee2 name:", employee2.name)
// }

// * struct : embedded struct
// type Person struct {
// 	name	string
// 	age		int
// }

// type Employee struct {
// 	division	string
// 	person 		Person
// }

// func main() {
// 	var employee1 = Employee{}

// 	employee1.person.name = "Pratama"
// 	employee1.person.age = 23
// 	employee1.division = "Frontend Engineering"

// 	fmt.Printf("%+v", employee1)
// }

// * struct : anonymous struct
// type Person struct {
// 	name	string
// 	age		int
// }

// func main() {
// 	var employee1 = struct {
// 		person		Person
// 		division	string
// 	}{}

// 	employee1.person = Person{name: "Pratama", age: 23}
// 	employee1.division = "Frontend Engineering"

// 	var employee2 = struct {
// 		person		Person
// 		division	string
// 	}{
// 		person:		Person{name: "Fadhila", age: 23},
// 		division: "Engineering",
// 	}

// 	fmt.Printf("Employee1: %+v\n", employee1)
// 	fmt.Printf("Employee2: %+v\n", employee2)
// }

// * struct : slice of struct
// type Person struct {
// 	name	string
// 	age		int
// }

// func main() {
// 	var people = []Person{
// 		{name: "Airell", age: 23},
// 		{name: "Ananda", age: 23},
// 		{name: "Mailo", age: 23},
// 	}

// 	for _, v := range people {
// 		fmt.Printf("%+v\n", v)
// 	}
// }

// * struct : slice of anonymous struct
type Person struct {
	name	string
	age		int
}

func main() {
	var employee = []struct {
		person		Person
		division	string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 23}, division: "Marketing"},
	}

	for _,v := range employee {
		fmt.Printf("%+v\n", v)
	}


}