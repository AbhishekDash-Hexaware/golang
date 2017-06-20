package main

import (
	"io"
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
	for i:=0;i<7;i++ {
		io.WriteString(w,"hello\n")
	}
	fmt.Printf("hello called\n")

}
func hey(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hey world!")
	fmt.Printf("hey called\n")
}



func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hey",hey)
	fmt.Println("server up and running")
	http.ListenAndServe(":8000", nil)

}
