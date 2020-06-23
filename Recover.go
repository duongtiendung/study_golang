package main

import "fmt"

//Hàm recover trả về giá trị được truyền cho hàm panic và không bị Side Effect. Nghĩa là nếu goroutine không bị panic, hàm recover sẽ trả về nil
// func recover() interface{}

func defFoo() {
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {

		fmt.Println("WOHA! Program is panicking with value", r)
	}

	fmt.Println("defFoo() done")
}

func normMain() {

	fmt.Println("normMain() started")

	defer defFoo() // defer defFoo call

	if 1 == 0 {
		panic(nil) // panic here
	}

	fmt.Println("normMain() done")
}

func main() {
	fmt.Println("main() started")

	normMain() // normal call

	fmt.Println("main() done")
}
