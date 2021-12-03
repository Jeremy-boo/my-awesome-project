package main

import (
	"fmt"
	"log"
	"net/http"
)

// HealthyHandler healthy handler
func HealthyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	http.HandleFunc("/api/v1/healthy", HealthyHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}


