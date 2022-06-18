package main

import (
	"booking-service/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Booking service is running")

	mux := http.NewServeMux()

	mux.Handle("/fetch", http.HandlerFunc(api.Fetch))
	mux.Handle("/book", http.HandlerFunc(api.Book))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
