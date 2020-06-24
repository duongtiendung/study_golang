package main

import "fmt"

//Hàm panic dùng để phát sinh lỗi (dừng luôn)
//Khi panic xảy ra, chương trình chấm dứt, in ra tham số được truyền cho panic, theo kèm với thông tin stack lỗi.

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Printf("asd")
	lastName := "Elon"
	fullName(nil, &lastName)
	fmt.Println("returned normally from main")
}
