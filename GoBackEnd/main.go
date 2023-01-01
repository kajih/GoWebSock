package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Main")
	http.HandleFunc("/", rootPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func rootPage(w http.ResponseWriter, _ *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "RESPONSE")
}
