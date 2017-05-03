package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
	for i:=0;i<7;i++ {
		io.WriteString(w,"hello\n")
	}




}
func hey(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hey world!")
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hey",hey)
	http.ListenAndServe(":8000", nil)

}
