package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id            uint    `json:"id"`
	Name          string  `json:"name"`
	Color         string  `json:"color"`
	Price         float64 `json:"price"`
	Stock         uint    `json:"stock"`
	Code          string  `json:"code"`
	Published     bool    `json:"published"`
	Creation_date string  `json:"creation_date"`
}

var products []product

func ReadFile() {
	jsonFile, err := os.Open(`products.json`)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &products)
	defer jsonFile.Close()
}

func GetAll(c *gin.Context) {
	ReadFile()
	c.JSON(200, products)
}

func GetById(c *gin.Context) {
	ReadFile()
	id := c.Param("id")
	for _, item := range products {
		if id == fmt.Sprint(item.Id) {
			c.JSON(200, item)
		}
	}
	c.JSON(404, gin.H{"message": "Product not found"})
}

func GetAllBy(c *gin.Context) {
	filteredProducts := []product{}
	ReadFile()
	var id int
	var err error
	if c.Query("id") != "" {
		id, err = strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid Id"})
			return
		}
	}

	name := c.Query("mame")
	color := c.Query("color")

	var price float64
	if c.Query("price") != "" {
		price, err = strconv.ParseFloat(c.Query("price"), 64)

		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price"})
			return
		}
	}
	var stock int
	if c.Query("stock") != "" {
		stock, err = strconv.Atoi(c.Query("stock"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid stock"})
			return
		}
	}

	code := c.Query("code")
	var published bool
	if c.Query("published") != "" {
		published, err = strconv.ParseBool(c.Query("published"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid published"})
			return
		}
	}
	creation_date := c.Query("creation_date")

	for _, product := range products {
		if (int(product.Id) == id || c.Query("id") == "") &&
			(product.Name == name || c.Query("name") == "") &&
			(product.Color == color || c.Query("color") == "") &&
			(product.Price == price || c.Query("price") == "") &&
			(int(product.Stock) == stock || c.Query("stock") == "") &&
			(product.Code == code || c.Query("code") == "") &&
			(product.Published == published || c.Query("published") == "") &&
			(product.Creation_date == creation_date || c.Query("createdAt") == "") {
			filteredProducts = append(filteredProducts, product)
		}
	}

	if len(filteredProducts) == 0 {
		c.JSON(404, gin.H{
			"error": "Produto não encontrado",
		})
		return
	}

	c.JSON(200, gin.H{
		"Produtos": filteredProducts,
	})
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Olá Davi",
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello", Hello)
	router.GET("/products/:id", GetById)
	router.GET("/products/all", GetAll)
	router.GET("/products", GetAllBy)
	router.Run()
}
