package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey Gopher, you've requested %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", HelloHandler)

	fmt.Println("Server listening to port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
