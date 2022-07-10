package pickbuy

import (
	"database/sql"
	"db_p/structs"
	"fmt"
)

func GetAllCategories(db *sql.DB) []structs.Categories {
	var categories []structs.Categories
	var category structs.Categories

	result, err := db.Query("SELECT distinct cat_id, cat_title FROM categories ")
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&category.Cat_id, &category.Cat_title)
		categories = append(categories, category)
		if err != nil {
			panic(err)
		}
	}
	return categories
}
func showAndChooseCategory(allCategories []structs.Categories) structs.Categories {
	println("Choose a category:")
	dictCategories := make(map[int]structs.Categories)
	var j int
	i := 0
	for i < len(allCategories) {
		dictCategories[i] = allCategories[i]
		//todo show more info
		fmt.Printf("%d ) %s\n", i+1, dictCategories[i].Cat_title)
		i += 1
	}
	fmt.Scanf("%d", &j)
	//	fmt.Printf("\n%d\n", dictCategories[j].Cat_id)
	return dictCategories[j-1]
}
func GetProductsByCat(category int, db *sql.DB) []structs.Product {
	var products []structs.Product

	rows, err := db.Query("SELECT * FROM products WHERE product_cat=?", category)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var p structs.Product

		err = rows.Scan(&p.Product_id, &p.Product_cat, &p.Product_brand, &p.Product_title,
			&p.Product_price, &p.Product_desc, &p.Product_image, &p.Product_keywords, &p.Product_count)

		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	return products

}
func showAndChooseProduct(products []structs.Product) structs.Product {

	println("\nchoose a product.")
	dictProducts := make(map[int]structs.Product)
	var j int

	i := 0
	for i < len(products) {
		dictProducts[i] = products[i]
		//todo show more info
		fmt.Printf("%d ) %s\n", i+1, dictProducts[i].Product_title)
		i += 1
	}
	fmt.Scanf("%d", &j)
	return dictProducts[j]

}
func selectProduct(product *structs.Product, chosenProducts *[]structs.Product, sumPrices *int) { //([]structs.Product ,int){
	//check if the product is available ?

	if product.Product_count >= 0 {
		*sumPrices += product.Product_price
		*chosenProducts = append(*chosenProducts, *product)
	} else {
		println("this product is unavailable for now.")
	}
	//	return chosenProducts, sumPrices
}
func showOrders(selectedProducts []structs.Product, totalAmount int) {
	println("product title           price   :\n")
	i := 0
	for i < len(selectedProducts) {
		p := selectedProducts[i]
		fmt.Printf("%d ) %s %d\n", i+1, p.Product_title, p.Product_price)
		i += 1
	}
	fmt.Printf("total sum :\t\t\t %d", totalAmount)

}
func Pick(db *sql.DB) ([]structs.Product, int) {
	//show categories
	var sumPrices = 0
	var chosenProducts []structs.Product
	allCategories := GetAllCategories(db)
	category := showAndChooseCategory(allCategories)
	fmt.Printf("%d", category.Cat_id)
	products := GetProductsByCat(category.Cat_id, db)
	var exit string
	for true {
		product := showAndChooseProduct(products)
		println("wanna select this product?\n1)Yse\n2)No\n")
		var ans int
		fmt.Scanf("%d", &ans)
		if ans == 1 {
			//chosenProducts,sumPrices=
			selectProduct(&product, &chosenProducts, &sumPrices)
		}
		println("to exit enter EXIT.") //yes of no
		fmt.Scanf("%s", &exit)
		if exit == "EXIT" {
			break
		}
	}
	showOrders(chosenProducts, sumPrices)
	return chosenProducts, sumPrices
}

func Buy(db *sql.DB, id int) {
	products, sumAmount := Pick(db)
	showOrders(products, sumAmount)

}
