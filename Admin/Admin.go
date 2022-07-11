func InsertNewProduct(db *sql.DB, p Product) error {  
    query := "INSERT INTO products(product_cat, product_brand , product_title , product_price ,product_desc ,product_image , product_keywords  ) 
    VALUES (?, ? , ? , ? , ? , ? , ?)"
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    stmt, err := db.PrepareContext(ctx, query)
    if err != nil {
    log.Printf("Error %s when preparing SQL statement", err)
    return err
    }
    defer stmt.Close()
}




func (db *sql.DB) Delete(id int) (int64, error) {
	result, err := db.Exec("delete from products where id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}






