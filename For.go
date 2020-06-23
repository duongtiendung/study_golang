package main

import (
	"fmt"
)

func main() {

	// for

	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	//while

	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)

	//foreach key & value

	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	//exit loop

	sum = 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum)

}
