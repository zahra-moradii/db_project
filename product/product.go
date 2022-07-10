package product

import (
	"database/sql"
	"fmt"
)

type Product struct {
	PID       int
	Name     string
	Category string
	Quantity int
	Price    float64
}

type Order struct {
	OID       string    `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Products []Product `json:"products"`
	Price    float64   `json:"price"`
	Status   string    `json:"status"`
}

type OrderedProduct struct {
	OPID              string
	ProductId       string
	ProductQuantity int
	OrderId         string
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error while reading all products from database: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.PID, &p.Name, &p.Category, &p.Quantity, &p.Price); err != nil {
			return nil, fmt.Errorf("parsing to a product failed with: %v", err)
		}
		products = append(products, p)
	}

	return products, nil
}

func GetProductById(productId string, db *sql.DB) (*Product, error) {
	row := db.QueryRow("SELECT * FROM products WHERE id = ?", productId)

	var p Product
	if err := row.Scan(&p.PID, &p.Name, &p.Category, &p.Quantity, &p.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no product with id: %s", productId)
		}
		return nil, fmt.Errorf("searching for %s failed with: %s", productId, err)
	}

	return &p, nil
}

func GetAllOrders(db *sql.DB) ([]Order, error) {
	var orders []Order

	rows, _ := db.Query("SELECT * FROM orders")
	defer rows.Close()

	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.OID, &o.Name, &o.Address, &o.Phone, &o.Price, &o.Status); err != nil {
			return nil, fmt.Errorf("getting all products failed with: %v", err)
		}

		products, err := GetAllProductsForOrder(o.OID)
		if err != nil {
			return nil, err
		}
		o.Products = products

		orders = append(orders, o)
	}

	return orders, nil
}

func GetOrderById(orderId string, db *sql.DB) (*Order, error) {
	row := db.QueryRow("SELECT * FROM orders WHERE id = ?", orderId)

	var o Order
	if err := row.Scan(&o.OID, &o.Name, &o.Address, &o.Phone, &o.Price, &o.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no order with id: %s", orderId)
		}
		return nil, fmt.Errorf("searching for %s failed with: %s", orderId, err)
	}

	products, err := GetAllProductsForOrder(o.OID)
	if err != nil {
		return nil, err
	}
	o.Products = products

	return &o, nil
}

func UpdateProduct(product *Product, db *sql.DB) error {

	result, err := db.Exec("UPDATE products SET NAME = ?, CATEGORY = ?, QUANTITY = ?, PRICE = ? WHERE ID = ?", product.Name, product.Category, product.Quantity, product.Price, product.PID)
	if err != nil {
		return fmt.Errorf("failed to update product to the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product with id: %s", product.PID)
	}

	return nil
}

func UpdateOrder(order *Order, db *sql.DB) error {

	result, err := db.Exec("UPDATE orders SET NAME = ?, ADDRESS = ?, PHONE = ?, PRICE = ? WHERE ID = ?", order.Name, order.Address, order.Phone, order.Price, order.OID)
	if err != nil {
		return fmt.Errorf("failed to update order to the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product with id: %s", order.OID)
	}

	return nil
}

func DeleteOrder(orderId string, db *sql.DB) error {
	result, err := db.Exec("DELETE FROM orders WHERE ID = ?;", orderId)
	if err != nil {
		return fmt.Errorf("failed to delete order from the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no order with id: %d", orderId)
	}

	return nil
}

func DeleteAllProductsForAnOrder(orderId string, db *sql.DB) error {
	result, err := db.Exec("DELETE FROM orderedProduct WHERE ORDER_ID = ?;", orderId)
	if err != nil {
		return fmt.Errorf("failed to delete ordered product from the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no order with id: %d", orderId)
	}

	return nil
}

func DeleteProduct(productId string, db *sql.DB) error {
	result, err := db.Exec("DELETE FROM products WHERE ID = ?;", productId)
	if err != nil {
		return fmt.Errorf("failed to delete product from the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product with id: %d", productId)
	}

	return nil
}

func ChangeProductQuantity(productId string, quantity int, db *sql.DB) error {
	var p Product

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", productId)
	if err := row.Scan(&p.PID, &p.Name, &p.Category, &p.Quantity, &p.Price); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no product with id: %s", productId)
		}
		return fmt.Errorf("searching for id: %s failed with: %s", productId, err)
	}

	newQuantity := p.Quantity - quantity
	if newQuantity < 0 {
		return fmt.Errorf("not enough quantity of product: %s", p.Name)
	}

	if _, err := db.Query("UPDATE products SET quantity = ? WHERE id = ?", newQuantity, p.PID); err != nil {
		return fmt.Errorf("updating quantity failed with: %s", err)
	}

	return nil
}

func GetAllProductsForOrder(orderId string, db *sql.DB) ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT product_id, quantity FROM orderedProduct WHERE order_id = ?", orderId)
	if err != nil {
		return nil, fmt.Errorf("error while reading ordered product from database: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.PID, &p.Quantity); err != nil {
			return nil, fmt.Errorf("parsing to a product failed with: %v", err)
		}

		details, err := GetProductById(p.PID)
		if err != nil {
			return nil, err
		}
		p.Name, p.Category, p.Price = details.Name, details.Category, details.Price

		products = append(products, p)
	}

	return products, nil
}
