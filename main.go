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

	r.GET("/profile/{id}/newProduct", API.Logs)
	r.GET("/profile/{id}/newProduct", API.News)
	r.PUT("profile/{id}/modifyName", API.ModifyFirstName)
	r.PUT("profile/{id}/modifyName", API.ModifyLastName)
	r.PUT("profile/{id}/modifyName", API.ModifyEmail)
	r.PUT("profile/{id}/modifyName", API.ModifyPassword)
	r.PUT("profile/{id}/modifyName", API.ModifyMobile)
	r.PUT("profile/{id}/modifyName", API.ModifyAdd1)
	r.PUT("profile/{id}/modifyName", API.ModifyAdd2)

	r.GET("/allCategories", API.AllCategories)
	r.GET("allCategories/:catId", API.ProductsByCategory)
	r.GET("/products", API.AllProducts)
	r.GET("/products/:productId", API.GetAllProductsByCategory)
	r.GET("/products/:productId/recommends", API.GetRecommend)

	r.GET("/order", API.GetAllOrders)
	/*
		//	r.GET("/product/:category", handler.GetProductCatHandler)
		r.GET("/product/:productId", handler.GetProductHandler)
		r.GET("/order/:orderId", handler.GetOrderHandler)

		r.POST("/order", handler.AddOrderHandler)
		r.POST("/product", handler.AddProductHandler)

		r.PUT("/order/:orderId", handler.UpdateOrderHandler)
		r.PUT("/product/:productId", handler.UpdateProductHandler)

		r.DELETE("/delete/product/:productId", handler.DeleteProductHandler)
		r.DELETE("/delete/order/:orderId", handler.DeleteOrderHandler)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	*/
	log.Println("Listening to port 8080...")
	log.Fatal(r.Run(":8080"))

	//id, _ := signUP_IN.SignIn(db)
	//pickbuy.Order(db, id)

	//pickbuy.Buy(db, id)
}
