package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

func main() {
	fmt.Println("Start database application...")
	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal("Occured fatal error: ", err.Error())
	}

	rows, err := db.Query("SELECT id, name, inventory, price FROM products")
	if err != nil {
		log.Fatal("Occured fatal error while querying: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		fmt.Printf("Row - %v \n", p)
	}
}
