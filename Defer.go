//Defer là một thứ gì đó mà luôn luôn được thực thi ở cuối một function.

package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func third() {
	fmt.Println("3rd")
}

// hàm first sẽ đc chạy sau cùng.

func main() {
	defer first()
	second()
	defer third()
}
