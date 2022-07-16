package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	. "github.com/golang-rest-shop-backend/pkg/structs"
)

var db *sql.DB

func InitMySqlConnection() error {

	config := mysql.Config{
		User: "root",
		//os.Getenv("MYSQL_USER"),
		Passwd: "m9m8m7m6@021",
		//os.Getenv("MYSQL_PASSWORD"),
		Net:  "tcp",
		Addr: "127.0.0.1:3306",
		//os.Getenv("MYSQL_IP_ADDRESS"),
		DBName:               "online_shop",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return fmt.Errorf("database opening failed with error: %s", err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return fmt.Errorf("ping my sql failed with error: %s", err.Error())
	}

	return nil
}

func GetAllProducts() ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT product_id , product_title , product_cat , product_count , product_price FROM products")
	if err != nil {
		return nil, fmt.Errorf("error while reading all products from database: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Product_id, &p.Product_title, &p.Product_cat, &p.Product_count, &p.Product_price); err != nil {
			return nil, fmt.Errorf("parsing to a product failed with: %v", err)
		}
		products = append(products, p)
	}

	return products, nil
}

func GetProductById(productId string) (*Product, error) {
	row := db.QueryRow("SELECT * FROM products WHERE product_id = ?", productId)

	var p Product
	if err := row.Scan(&p.Product_id, &p.Product_title, &p.Product_cat, &p.Product_count, &p.Product_price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no product with id: %s", productId)
		}
		return nil, fmt.Errorf("searching for %s failed with: %s", productId, err)
	}

	return &p, nil
}

func AddProduct(product *Product) (string, error) {

	err, _ := db.Query("INSERT INTO products (Product_id, Product_cat, "+
		"Product_brand, Product_title, Product_price , Product_desc , Product_image , Product_keywords , "+
		"Product_count) VALUES (?,?,?,?,?,?,?,?,? )", product.Product_id,
		product.Product_cat, product.Product_brand, product.Product_title, product.Product_price,
		product.Product_desc, product.Product_image, product.Product_keywords, product.Product_count)
	if err != nil {
		return "", fmt.Errorf("failed to add product to the database, error: %s", err)
	}

	return string(rune(product.Product_id)), nil
}

func UpdateProduct(product *Product) error {

	result, err := db.Exec("UPDATE products SET product_title = ?, product_cat = ?, product_count = ?, product_price = ? WHERE product_id = ?", product.Product_title, product.Product_cat, product.Product_count, product.Product_price, product.Product_id)
	if err != nil {
		return fmt.Errorf("failed to update product to the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product with id: %s", product.Product_id)
	}

	return nil
}

func DeleteProduct(productId string) error {
	result, err := db.Exec("DELETE FROM products WHERE product_id = ?;", productId)
	if err != nil {
		return fmt.Errorf("failed to delete product from the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product with id: %d", productId)
	}

	return nil
}

func ChangeProductQuantity(productId string, quantity int) error {
	var p Product

	row := db.QueryRow("SELECT * FROM products WHERE product_id = ?", productId)
	if err := row.Scan(&p.Product_id, &p.Product_title, &p.Product_cat, &p.Product_count, &p.Product_price); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no product with id: %s", productId)
		}
		return fmt.Errorf("searching for id: %s failed with: %s", productId, err)
	}

	newQuantity := p.Product_count - quantity
	if newQuantity < 0 {
		return fmt.Errorf("not enough quantity of product: %s", p.Product_title)
	}

	if _, err := db.Query("UPDATE products SET product_count = ? WHERE product_id = ?", newQuantity, p.Product_id); err != nil {
		return fmt.Errorf("updating quantity failed with: %s", err)
	}

	return nil
}
