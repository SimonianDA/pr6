package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Title string `json:"title"`
}

var products = []product{
	{ID: "1", Type: "Computer", Title: "Storage devices"},
	{ID: "2", Type: "Computer", Title: "RAM"},
	{ID: "3", Type: "Computer", Title: "CPU"},
	{ID: "4", Type: "Computer", Title: "Motherboard"},
	{ID: "5", Type: "Computer", Title: "Cooler"},
	{ID: "6", Type: "Computer", Title: "Case"},
	{ID: "7", Type: "Computer", Title: "Keyboard"},
	{ID: "8", Type: "Computer", Title: "Mouse"},
	{ID: "9", Type: "Computer", Title: "Headphones"},
	{ID: "10", Type: "Computer", Title: "Monitor"},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	router.Run(":8080")
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func createProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct product

	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
