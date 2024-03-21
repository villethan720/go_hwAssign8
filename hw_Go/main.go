package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:4000", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./superHero.txt")
	io.Copy(w, f)
}
