package main

import (
	"fmt"
)

func main() {
	fmt.Println("go stat")

	//If statement
	a := 1
	fmt.Printf("hãy nhập a \n")
	fmt.Scanf("%d", &a)

	fmt.Println(a)

	if a > 1 {
		fmt.Printf(" a>1 \n")
	} else {
		fmt.Printf(" a=<1 \n")
	}

	b := 1
	fmt.Printf("hãy nhập b \n")
	fmt.Scanf("%d", &b)
	//For statement
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	fmt.Println(sum)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hi")
	// })

	// log.Fatal(http.ListenAndServe(":8081", nil))

}
