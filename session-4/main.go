package main

import (
	"fmt"
	"sync"
)

// * interface
// type shape interface {
// 	area() float64
// 	perimeter() float64
// }

// type rectangle struct {
// 	width, height float64
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * math.Pow(c.radius, 2)
// }

// func (r rectangle) area() float64 {
// 	return r.height * r.width
// }

// func (c circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func (r rectangle) perimeter() float64 {
// 	return 2 * (r.height + r.width)
// }

func main() {
	// var c1 shape = circle{radius: 5}
	// var r1 shape = rectangle{width: 3, height: 2}

	// fmt.Printf("Type of c1: %T\n", c1)
	// fmt.Printf("Type of r1: %T\n", r1)

	// fmt.Println("Circle area", c1.area())
	// fmt.Println("Circle perimeter", c1.perimeter())

	// fmt.Println("Rectangle area", r1.area())
	// fmt.Println("Rectangle perimeter", r1.perimeter())
	// print("Rectangle", r1)
	// print("Circle", c1)

	// c1.volume()
	// * interface : type assertion
	// c1.(circle).volume()

	// value, ok := c1.(circle)

	// if ok {
	// 	fmt.Printf("Circle value: %+v\n", value)
	// 	fmt.Printf("Circle volume: %+f\n", value.volume())
	// }

	// * empty interface
	// var randomValues interface{}

	// _ = randomValues

	// randomValues = "Jalan Sudirman"
	// randomValues = 20
	// randomValues = true
	// randomValues = []string{"Airell", "Nanda"}

	// * empty interface : type assertion
	// var v interface{}
	// v = 20
	// // v = v * 9 --> ini akan error karena v bertipe interface tidak bisa dikalikan dengan 9 tipe int asli
	// if value, ok := v.(int); ok {
	// 	v = value * 9
	// }

	// * empty interface : map & slices with empty interface
	// rs := []interface{}{1, "Airell", true, 2, "Ananda", true}
	// rm := map[string]interface{}{
	// 	"Name":   "Airell",
	// 	"Status": true,
	// 	"Age":    23,
	// }
	// _, _ = rs, rm

	// * reflect : identifying data type & value
	// var number = 23
	// var reflectValue = reflect.ValueOf(number)

	// fmt.Println("Tipe variabel :", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("Nilai variabel :", reflectValue.Int())
	// }

	// * reflect : accessing value using interface method
	// fmt.Println("Nilai variabel (menggunakan interface) :", reflectValue.Interface())

	// var nilai = reflectValue.Interface().(int)
	// fmt.Println(nilai)

	// * reflect : identifying method information
	// var s1 = &student{Name: "john wick", Grade: 2}
	// fmt.Println("nama :", s1.Name)

	// var reflectValueS1 = reflect.ValueOf(s1)
	// var method = reflectValueS1.MethodByName("SetName")
	// method.Call([]reflect.Value{
	// 	reflect.ValueOf("wick"),
	// })

	// fmt.Println("nama :", s1.Name)

	// * concurrency - goroutines
	// go goroutine()

	// * goroutines : asynchronous process
	// fmt.Println("main execution started")

	// go firstProcess(8)
	// secondProcess(8)

	// fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	// time.Sleep(time.Second * 2)
	// fmt.Println("main execution ended")

	// * waitgroup for synchronize goroutine
	fruits := []string{"apple", "mango", "durian", "rambutan"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}

	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index ==> %d, fruit ==> %s\n", index, fruit)
	wg.Done()
}

// func firstProcess(index int) {
// 	fmt.Println("First process func started")
// 	for i := 1; i <= index; i++ {
// 		fmt.Println("i=", i)
// 	}
// 	fmt.Println("First process func ended")
// }

// func secondProcess(index int) {
// 	fmt.Println("Second process func started")
// 	for j := 1; j <= index; j++ {
// 		fmt.Println("j=", j)
// 	}
// 	fmt.Println("Second process func ended")
// }

// func goroutine() {
// 	fmt.Println("Hello")
// }

// func print(t string, s shape) {
// 	fmt.Printf("%s area %v\n", t, s.area())
// 	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
// }

// func (c circle) volume() float64 {
// 	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
// }

// type student struct {
// 	Name  string
// 	Grade int
// }

// func (s *student) SetName(name string) {
// 	s.Name = name
// }
