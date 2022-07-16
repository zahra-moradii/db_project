package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-rest-shop-backend/pkg/service"
	"github.com/golang-rest-shop-backend/pkg/structs"
	"net/http"
	"strconv"
	_ "strconv"
)

// @Router /product [get]
func GetAllProductHandler(c *gin.Context) {
	currency := c.Param("currency")

	products, err := service.GetAllProducts(currency)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Router /product/{productId} [get]
func GetProductHandler(c *gin.Context) {
	currency := c.Param("currency")
	productId := c.Param("productId")

	product, err := service.GetProductById(productId, currency)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())

		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Router /product [post]
func AddProductHandler(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var product structs.Product
	err := decoder.Decode(&product)
	if err != nil {
		c.String(http.StatusBadRequest, "request body has wrong format: %s\n", err)

		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	productID, err := service.AddProduct(&product)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.String(http.StatusOK, "Product successfully added id: %s", productID)
}

// @Router /product/{productId} [put]
func UpdateProductHandler(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var product structs.Product
	err := decoder.Decode(&product)
	if err != nil {
		c.String(http.StatusBadRequest, "request body has wrong format: %s\n", err)

		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	product.Product_id, _ = strconv.Atoi(c.Param("productId"))
	product.Product_desc = c.Param("productdesc")
	product.Product_image = c.Param("productImage")
	product.Product_price, _ = strconv.Atoi(c.Param("productId"))

	if err = service.UpdateProduct(&product); err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusInternalServerError, "Product %s updated succesfully!", product.Product_id)
}

// @Router /delete/product/{productId} [delete]
func DeleteProductHandler(c *gin.Context) {
	productId := c.Param("productId")
	productPrice := c.Param("productPrice")

	if err := service.DeleteProduct(productId); err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "Product %s deleted %s ", productPrice)
}
