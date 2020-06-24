package main

import "fmt"

//Hàm recover trả về giá trị được truyền cho hàm panic và không bị Side Effect. Nghĩa là nếu goroutine không bị panic, hàm recover sẽ trả về nil
// func recover() interface{}

func defFoo(itf interface{}) {
	fmt.Printf("Type = %T, value = %v\n", itf, itf)
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {

		fmt.Println("WOHA! Program is panicking with value", r)
	}

	fmt.Println("defFoo() done")
}

func normMain() {

	fmt.Println("normMain() started")

	defer defFoo("viet truoc panic") // defer defFoo call

	if 1 == 1 {
		panic("huhu") // panic here
	}

	defer defFoo("viet sau panic") // defer defFoo call

	fmt.Println("normMain() done")
}

func main() {
	fmt.Println("main() started")

	normMain() // normal call

	fmt.Println("main() done")
}
