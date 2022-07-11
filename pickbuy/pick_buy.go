package pickbuy

import (
	"database/sql"
	"db_p/profile"
	"db_p/structs"
	"fmt"
	"time"
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
func selectProduct(product *structs.Product, chosenProducts *[]Pair, sumPrices *int) {
	//check if the product is available ?
	var amount int
	println("how many?")
	fmt.Scanf("%d", &amount)

	if product.Product_count*amount >= 0 {
		*sumPrices += product.Product_price * amount
		*chosenProducts = append(*chosenProducts, Pair{*product, amount})
	} else {
		//todo inform the costumer when it became available
		println("this product is unavailable for now.\n")
	}

}
func showOrders(selectedProducts []Pair, totalAmount int) {
	println("product title\t\tprice\t\tamount\n")
	i := 0
	for i < len(selectedProducts) {
		p := selectedProducts[i]
		fmt.Printf("%d ) %s %d \t\t%d\n", i+1, p.product.Product_title, p.product.Product_price, p.amount)
		i += 1
	}
	fmt.Printf("total sum :\t\t %d\n", totalAmount)

}

type Pair struct {
	product structs.Product
	amount  int
}

func Pick(db *sql.DB) ([]Pair, int) {
	//show categories
	var sumPrices = 0
	var chosenProducts []Pair
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
	return chosenProducts, sumPrices
}
func Order(db *sql.DB, id int) {
	products, sumAmount := Pick(db)
	showOrders(products, sumAmount)
	i := 0
	for i < len(products) {
		p := products[i]
		_, err := db.Query(`INSERT INTO orders  SET user_id=?,product_id=?,total_amt=?,count=?,status=?  `,
			id, p.product.Product_id, p.amount*p.product.Product_price, p.amount, "SELECTED")

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		i += 1
	}
}
func chooseAddress(db *sql.DB, id int) string {
	var address [3]string
	result, err := db.Query(`SELECT address1,address2 FROM user_info WHERE user_id=?`, id)
	for result.Next() {
		err = result.Scan(&address[1], &address[2])
		if err != nil {
			panic(err)

		}
	}
	//choose address
	add := -1
	if address[1] != "X" {
		fmt.Printf("%d)address1:%s", 1, address[1])
		if address[2] != "X" {
			fmt.Printf("%d)address2:%s", 2, address[2])
			fmt.Scanf("%d", &add)
		}

	} else {
		println("please enter your address:")
		//todo modify address
		fmt.Scanf("%s", &address[0])
		profile.Modify_address1(id, address[0], db)

	}
	if add != -1 {
		address[0] = address[add]
	}
	return address[0]
}

func GetProductById(db *sql.DB, product_id int, p *structs.Product) {
	rows, err := db.Query("SELECT * FROM products WHERE product_id=?", product_id)
	if err != nil {
		panic(err)
	}

	for rows.Next() {

		err = rows.Scan(&p.Product_id, &p.Product_cat, &p.Product_brand, &p.Product_title,
			&p.Product_price, &p.Product_desc, &p.Product_image, &p.Product_keywords, &p.Product_count)

		if err != nil {
			panic(err)
		}
	}

}
func updateLog(db *sql.DB, id int, orderID int, totoalAmount int, cardNum int, cvv int, address string) {

	//make a log
	date := time.Now().Format("2006.01.02 15:04:05")
	_, err := db.Query(`INSERT INTO logs  SET order_id=?,user_id=?,action=?,address=?,total_amt=?,card_number=?,cvv=?,date=?`,
		orderID, id, "COMPLETED", address, totoalAmount, cardNum, cvv, date)
	if err != nil {
		panic(err.Error())
	}

	//update status

	_, err = db.Query(`UPDATE orders SET   status= ?	WHERE  order_id=? and user_id=?`,
		"COMPLETED", orderID, id)

	if err != nil {
		panic(err.Error())
	}

}
func Buy(db *sql.DB, id int) {
	var count int
	var orderTotalCost int
	var orderID int
	var product_id int
	var arrOrderIDs []int
	sum := 0
	var products []Pair
	var product structs.Product
	//find selected products
	result, err := db.Query(`SELECT order_id,product_id,count,total_amt FROM orders WHERE status=? and user_id=? `, "SELECTED", id)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for result.Next() {
		err = result.Scan(&orderID, &product_id, &count, &orderTotalCost)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		arrOrderIDs = append(arrOrderIDs, orderID)
		sum += orderTotalCost
		GetProductById(db, product_id, &product)
		products = append(products, Pair{product, count})
	}
	//show products
	println("all your orders:")
	showOrders(products, sum)
	var ans int
	println("\nAre you sure you want to buy?\n0)NO\n1)YES\n")
	fmt.Scanf("%d", &ans)
	if ans == 1 {
		//log
		var cardNum int
		var cvv int

		println("enter your card number:")
		fmt.Scanf("%d", &cardNum)

		println("enter your cvv number:")
		fmt.Scanf("%d", &cvv)

		address := chooseAddress(db, id)

		i := 0
		for i < len(arrOrderIDs) {
			//i is order id
			updateLog(db, id, arrOrderIDs[i], sum, cardNum, cvv, address)
			i += 1
		}
	} else {
		return
	}

}
