package product

import (
	"database/sql"
	"errors"
	"fmt"
)


type Product struct {
	ID       int
	Name     string
	Category string
	Quantity int
	Price    float64
}


func GetAllProducts() ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error while reading all products from database: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Quantity, &p.Price); err != nil {
			return nil, fmt.Errorf("parsing to a product failed with: %v", err)
		}
		products = append(products, p)
	}

	return products, nil
}