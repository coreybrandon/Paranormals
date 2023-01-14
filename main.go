package main

import (
	"fmt"
	"net/http"

	"github.com/coreybrandon/haunting/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/products", handlers.GetProducts)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)

	}
}
