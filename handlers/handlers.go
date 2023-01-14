package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coreybrandon/haunting/models"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test\n")
}

func GetProducts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var products []models.Product
	json.NewEncoder(w).Encode(products)

}
