package main

import (
	"fmt"
	"time"
)

func main() {

	// for

	// sum := 0
	// for i := 1; i < 5; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// //while

	// n := 1
	// for n < 5 {
	// 	n *= 2
	// }
	// fmt.Println(n)

	//foreach key & value

	strings := []string{"hello", "world", "a", "b", "c"}
	fmt.Println(" in array lần 1 ============================")
	for i, s := range strings {
		fmt.Println(i, s)
	}
	fmt.Println(" in array lần 2 ============================")
	for i, s := range strings {
		fmt.Println(i, s)
	}

	maps := map[int]string{
		0: "hello",
		1: "world",
		2: "a",
		3: "b",
		4: "c",
	}
	fmt.Println(" in map lần 1 ============================")
	for i, s := range maps {
		time.Sleep(1 * time.Second)
		fmt.Println(i, s)
	}
	fmt.Println("in map lần 2 ============================")
	for i, s := range maps {
		time.Sleep(1 * time.Second)
		fmt.Println(i, s)
	}

	//exit loop

	// sum = 0
	// for i := 1; i < 5; i++ {
	// 	if i%2 != 0 { // skip odd numbers
	// 		continue
	// 	}
	// 	sum += i
	// }
	// fmt.Println(sum)

}
