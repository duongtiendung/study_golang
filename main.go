package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	b := 1
	//Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Printf("Nhấn Enter để tiếp tục \n")
	fmt.Scanf("%d", &b)
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Printf("Nhấn Enter để tiếp tục \n")
	fmt.Scanf("%d", &b)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("Nhấn Enter để tiếp tục \n")
	fmt.Scanf("%d", &b)
	fmt.Println(primes)
	//

	fmt.Printf("Nhấn Enter để tiếp tục \n")
	fmt.Scanf("%d", &b)
	// map

	type Vertex struct {
		Lat, Long float64
	}
	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	//
	fmt.Printf("Nhấn Enter để tiếp tục \n")
	fmt.Scanf("%d", &b)
	//Pointers

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	//
	fmt.Printf("Nhấn Enter để tiếp tục \n")
	fmt.Scanf("%d", &b)
	//Structs
	type Vertex1 struct {
		X int
		Y int
	}

	fmt.Printf("Nhấn Enter để tiếp tục \n")
	fmt.Scanf("%d", &b)
	// Định nghĩa một struct với từ khóa type
	type Student struct {
		name string
		age  int
	}
	// Khởi tạo biến s1 có giá trị là struct Student
	s1 := Student{"Robin", 30} // {"Robin", 30}

	// Khởi tạo biến s2 có giá trị là struct Student với 1 field là name
	// Field còn lại sẽ có giá trị mặc định (zero value)
	s2 := Student{name: "Robin"} // {"Robin", 0}

	// Khởi tạo biến s3 có giá trị là struct Student và không khai báo giá trị cho trường nào
	s3 := Student{} // {"", 0}

	// Thay đổi giá trị field trong struct
	s3.name = "Robert"
	s3.age = 25

	fmt.Println(s3) // s3 = {"Robert", 25}

	s1 = Student{"Steve", 30}
	s2 = Student{"Steve", 30}
	s3 = Student{"Job", 30}

	if s1 == s2 {
		fmt.Println("s1 = s2t") // s1 bằng s2
	} else {
		fmt.Println("s1 != s2")
	}

	if s2 == s3 {
		fmt.Println("s2 = s3")
	} else {
		fmt.Println("s2 != s3") // s2 khác s3
	}
	fmt.Printf("Kết thúc phần code demo 3 \n")
	fmt.Scanf("%d", &b)

	// //If statement
	// d := 1
	// fmt.Printf("hãy nhập a \n")
	// fmt.Scanf("%d", &a)

	// fmt.Println(a)

	// if d > 1 {
	// 	fmt.Printf(" d>1 \n")
	// } else {
	// 	fmt.Printf(" d=<1 \n")
	// }

	// e := 1
	// fmt.Printf("hãy nhập b \n")
	// fmt.Scanf("%d", &e)
	// //For statement
	// sum := 0
	// for i := 0; i < 10000; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
