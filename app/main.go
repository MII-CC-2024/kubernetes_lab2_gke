package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
	} else {
		fmt.Fprintln(w, "<h1>Hello from", hostname, "(version 2022)</h1>")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Go Hello is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
