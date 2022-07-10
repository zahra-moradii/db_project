package pickbuy

import (
	"database/sql"
	"db_p/structs"
	"fmt"
)

func GetAllCategories(db *sql.DB) []string {
	var categories []string
	var category string
	result, err := db.Query("SELECT cat_title FROM categories ")
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&category)
		categories = append(categories, category)
		if err != nil {
			panic(err)
		}
	}
	return categories
}
func showCategory(allCategories []string) string {
	println("Choose a category:")
	//todo show categories and get a category and return it
	//for i := 0; i < len(allCategories); i++ {
	//if allCategories[i]
	//}
	return "elec" //allCategories[1] //todo
}
func GetProductsByCat(category string, db *sql.DB) []structs.Product {
	var products []structs.Product

	rows, err := db.Query("SELECT * FROM products WHERE product_cat=?", category)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var p structs.Product

		//todo image and ...
		err = rows.Scan(&p.ID, &p.Name, &p.Category, &p.Quantity, &p.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	return products

}
func showAndChooseProduct(products []structs.Product) structs.Product {
	//todo show products
}
func Pick(db *sql.DB, id int) []structs.Product {
	//show categories
	var chosenProducts []structs.Product
	allCategories := GetAllCategories(db)
	category := showCategory(allCategories)
	products := GetProductsByCat(category, db)
	var exit string
	for true {
		product := showAndChooseProduct(products)
		println("wanna pick this product?") //yes of no
		var ans int
		fmt.Scanf("%d", &ans)
		if ans == 1 {
			chosenProducts = append(chosenProducts, product)
		}
		println("to exit enter EXIT.") //yes of no
		fmt.Scanf("%s", &exit)
		if exit == "EXIT" {
			break
		}
	}
	return chosenProducts
}
