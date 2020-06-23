//switch
package main

import (
	"fmt"
	"time"
)

//Foo prints and returns n.
func Foo(n int) int {
	fmt.Println(n)
	return n
}

func main() {

	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ': // skip space
			break
		case '\n': // break at newline, quit Loop
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}

	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}
}
