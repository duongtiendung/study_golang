package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("go stat")

	//If statement
	a := 1
	fmt.Printf("hãy nhập a /n")
	fmt.Scanf("%f", &a)

	if a > 1 {
		fmt.Printf("Hiiiii... true")
	} else {
		fmt.Printf("Hiiiii... fales")
	}

	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	fmt.Println(sum)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
