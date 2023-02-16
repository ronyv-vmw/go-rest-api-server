package backend

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	Name        string `json:"name"`
	Inventory   int    `json:"inventory"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
}

func getProducts(db *sql.DB) ([]product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal("backend: Occured fatal error while querying: ", err.Error())
	}
	defer rows.Close()

	products := []product{}
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.ProductCode, &p.Name, &p.Inventory, &p.Price, &p.Status); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, err
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT id, productCode, status, name, inventory, price FROM products WHERE id = ?", p.ID).Scan(&p.ID, &p.ProductCode, &p.Status, &p.Name, &p.Inventory, &p.Price)
}
