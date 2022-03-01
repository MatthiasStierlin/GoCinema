package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//Get customer
	r.HandleFunc("/customer/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		intId, err := strconv.Atoi(vars["id"])

		fmt.Println(intId, err, reflect.TypeOf(intId))
		customer := Customer{}
		customer = getCustomer(intId)
		num := strconv.Itoa(customer.CustomerId)

		fmt.Fprintf(w, "The name of the customer: %s is %s\n", num, customer.Customername)
	}).Methods("GET")

	//Add customer
	r.HandleFunc("/customer/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		customer := Customer{}

		addCustomer(customer)
	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}
