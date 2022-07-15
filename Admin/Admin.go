package Admin

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func InsertNewProduct(db *sql.DB) error {
	query := "INSERT INTO products(product_cat, product_brand , product_title , product_price " +
		",product_desc ,product_image , product_keywords  )VALUES (?, ? , ? , ? , ? , ? , ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	return nil
}

//func Delete(id int, db *sql.DB) (int64, error) {
//	result, err := db.Exec("delete from products where id = ?", id)
//	if err != nil {
//		return 0, err
//	} else {
//		return result.RowsAffected(), nil
//	}
//}

func DeleteProduct(productId int, db *sql.DB) error {
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

func UpdateProduct(product *Product, db *sql.DB) error {

	result, err := db.Exec("UPDATE product SET NAME = ?, CATEGORY = ?, QUANTITY = ?, PRICE = ? WHERE ID = ?",
		product.Name, product.Category, product.Quantity, product.Price, product.ID)
	if err != nil {
		return fmt.Errorf("failed to update product to the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product with id: %s", product.ID)
	}
	return nil
}

func UpdateOrder(order *Orders, db *sql.DB) error {

	result, err := db.Exec("UPDATE orders SET User_id = ?, Product_id = ?, Qty = ? , Trx_id = ? , P_status = ? WHERE Order_id = ?",
		order.User_id, order.Product_id, order.Qty, order.Trx_id, order.P_status, order.Order_id)
	if err != nil {
		return fmt.Errorf("failed to update order to the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product with id: %s", order.Order_id)
	}

	return nil
}

func DeleteOrder(orderId int, db *sql.DB) error {
	result, err := db.Exec("DELETE FROM orders WHERE Order_id = ?;", orderId)
	if err != nil {
		return fmt.Errorf("failed to delete order from the database, error: %s", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no order with id: %d", orderId)
	}

	return nil
}

func ChangeProductQuantity(productId int, quantity int, db *sql.DB) error {
	var p Product

	row := db.QueryRow("SELECT * FROM products WHERE id = ?", productId)
	if err := row.Scan(&p.ID, &p.Name, &p.Category, &p.Quantity, &p.Price); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no product with id: %s", productId)
		}
		return fmt.Errorf("searching for id: %s failed with: %s", productId, err)
	}

	newQuantity := p.Quantity - quantity
	if newQuantity < 0 {
		return fmt.Errorf("not enough quantity of product: %s", p.Name)
	}

	if _, err := db.Query("UPDATE products SET quantity = ? WHERE id = ?", newQuantity, p.ID); err != nil {
		return fmt.Errorf("updating quantity failed with: %s", err)
	}

	return nil
}
