package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Customer struct {
	ID       string           `json:"id"`
	Customer *CustomerDetails `json:"customer"`
	Balance  uint             `json:"balance"`
}

type CustomerDetails struct {
	AccountNumber int    `json:"account_number"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Sex           string `json:"sex"`
}

var customers []Customer

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, each := range customers {
		if each.ID == params["id"] {
			json.NewEncoder(w).Encode(each)
			return
		}
	}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer Customer

	_ = json.NewDecoder(r.Body).Decode(&customer)
	customer.ID = strconv.Itoa(rand.Intn(10000000))
	customers = append(customers, customer)

	json.NewEncoder(w).Encode(customer)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, each := range customers {
		if each.ID == params["id"] {
			customers = append(customers[:index], customers[index+1:]...)

			var customer Customer
			_ = json.NewDecoder(r.Body).Decode(&customer)
			customer.ID = params["id"]
			customers = append(customers, customer)
			json.NewEncoder(w).Encode(customer)

			return
		}
	}
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, each := range customers {
		if each.ID == params["id"] {
			customers = append(customers[:index], customers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(customers)
}

func main() {
	r := mux.NewRouter()

	customers = append(
		customers,
		Customer{
			ID: "1",
			Customer: &CustomerDetails{
				AccountNumber: 361214402112,
				FirstName:     "Dafa",
				LastName:      "Armanto",
				Sex:           "Laki-laki",
			},
			Balance: 127.00,
		},
	)

	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", createAccount).Methods("POST")
	r.HandleFunc("/customers/{id}", updateAccount).Methods("PUT")
	r.HandleFunc("/customers/{id}", deleteAccount).Methods("DELETE")

	PORT := 8080
	fmt.Printf("Server start at port: %v", PORT)
	log.Fatalln(http.ListenAndServe(":8080", r))
}