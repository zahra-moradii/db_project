package main

import (
	API "db_p/Api"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	r := gin.Default()
	r.POST("/signUp", API.CreatUser)
	r.GET("/signIn", API.SignInUser)

	r.GET("/profile/:id/logs", API.Logs)
	r.PUT("/profile/:id/newProduct", API.News)
	r.PUT("profile/:id/modifyFName", API.ModifyFirstName)
	r.PUT("profile/:id/modifyLName", API.ModifyLastName)
	r.PUT("profile/:id/modifyEmail", API.ModifyEmail)
	r.PUT("profile/:id/modifyPass", API.ModifyPassword)
	r.PUT("profile/:id/modifyPhone", API.ModifyMobile)
	r.PUT("profile/:id/modifyAdd1", API.ModifyAdd1)
	r.PUT("profile/:id/modifyAdd2", API.ModifyAdd2)

	r.GET("/allCategories", API.AllCategories)
	r.GET("allCategories/:catId", API.ProductsByCategory)
	r.GET("/products", API.AllProducts)
	r.GET("/products/:productId", API.GetAllProductsByCategory)
	r.GET("/products/:productId/recommends", API.GetRecommend)

	r.GET("/order", API.GetAllOrders)
	r.POST("/order", API.AddOrder)
	r.DELETE("/delete/order/:orderId", API.DeleteOrder)

	r.PUT("/buyOrders", API.BuyOrders)

	log.Println("Listening to port 8080...")
	log.Fatal(r.Run(":8080"))

	//id, _ := signUP_IN.SignIn(db)
	//pickbuy.Order(db, id)

	//pickbuy.Buy(db, id)
}
