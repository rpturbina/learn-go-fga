package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	// ChannelImpl()
	// DeferAndExitImpl()
	// ErrorPanicRecoverImpl()

}

func ChannelImpl() {
	// * channel
	// c := make(chan string)

	// * channel operators
	// Mengirim data kepada channel
	// c <- value

	// Menerima data dari channel
	// result := <-c

	// * channel implementation
	// c := make(chan string)

	// go introduce("Airell", c)

	// go introduce("Nanda", c)

	// go introduce("Mailo", c)

	// msg1 := <-c
	// fmt.Println(msg1)

	// msg2 := <-c
	// fmt.Println(msg2)

	// msg3 := <-c
	// fmt.Println(msg3)

	// close(c)

	// * channel with anonymous function
	// c := make(chan string)
	// students := []string{"Airell", "Mailo", "Indah"}

	// for _, v := range students {
	// 	go func(student string) {
	// 		fmt.Println("Student", student)
	// 		result := fmt.Sprintf("Hai, my name is %s", student)
	// 		c <- result
	// 	}(v)
	// }

	// for i := 1; i <= 3; i++ {
	// 	print(c)
	// }

	// close(c)

	// * channel directions
	// c := make(chan string)

	// students := []string{"Airell", "Mailo", "Indah"}

	// for _, v := range students {
	// 	go introduce(v, c)
	// }

	// for i := 1; i <= 3; i++ {
	// 	print(c)
	// }

	// close(c)

	// * unbuffered channel
	// c1 := make(chan int)

	// go func(c chan int) {
	// 	fmt.Println("func goroutine starts sending data into the channel")
	// 	c <- 10
	// 	fmt.Println("func goroutine after sending data into the channel")
	// }(c1)

	// fmt.Println("main goroutine sleeps for 2 second")
	// time.Sleep(time.Second * 2)

	// fmt.Println("main goroutine starts receiving data")
	// d := <-c1
	// fmt.Println("main goroutine received data:", d)

	// close(c1)
	// time.Sleep(time.Second)

	// * buffered channel
	// c1 := make(chan int, 3)

	// go func(c chan int) {
	// 	for i := 1; i <= 5; i++ {
	// 		fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
	// 		c <- i
	// 		fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
	// 	}

	// 	close(c)
	// }(c1)

	// fmt.Println("main goroutine sleeps for 2 seconds")
	// time.Sleep(time.Second * 2)

	// for v := range c1 { // v= <- c1
	// 	fmt.Println("main goroutine received value from channel:", v)
	// }

	// * channel (select)
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c2 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}

func print(c <-chan string) {
	fmt.Println(<-c)
}

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}

func DeferAndExitImpl() {
	// * defer
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hei everyone")
	fmt.Println("Welcome back to Go learning center")

	callDeferFunc()
	fmt.Println("Hai everyone!!")

	// * exit
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

func ErrorPanicRecoverImpl() {
	// * error
	// var number int
	// var err error

	// number, err = strconv.Atoi("123GH")

	// if err == nil {
	// 	fmt.Println(number)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// number, err = strconv.Atoi("123")
	// if err == nil {
	// 	fmt.Println(number)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// * custom error
	// var password string

	// fmt.Scanln(&password)

	// if valid, err := validPassword(password); err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(valid)
	// }

	// * panic
	// var password string

	// fmt.Scanln(&password)

	// if valid, err := validPassword(password); err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Println(valid)
	// }

	// * recover
	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "Valid Password", nil
}
