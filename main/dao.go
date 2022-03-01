package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func getCustomer(id int) Customer {
	query := `SELECT c.customerId, c.customername FROM customers c WHERE c.customerId = ? LIMIT 1`

	db := initializeCon()
	row := db.QueryRow(query, id)
	db.Close()

	customer := Customer{}

	if err := row.Scan(&customer.CustomerId, &customer.Customername); err != nil {
		log.Fatalf("Could not scan row: %v", err)
	}

	fmt.Printf("Found customer: %+v\n", customer)
	return customer
}

/**
func addCustomer(customer Customer) {

}*/

func initializeCon() *sql.DB {
	db, err := sql.Open("mysql", "username:password@(localhost:port)/databasename")

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}

	fmt.Println("Database is reachable")

	return db
}
