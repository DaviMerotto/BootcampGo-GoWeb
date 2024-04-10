package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davimerotto/web-server/cmd/server/handler"
	"github.com/davimerotto/web-server/internal/products"
	"github.com/davimerotto/web-server/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TokenMiddleware(ctx *gin.Context) {
	tokenEnvironment := os.Getenv("TOKEN")
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token vazio"})
	}
	if token != tokenEnvironment {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
		return
	}
	ctx.Next()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	user := os.Getenv("MY_USER")
	password := os.Getenv("MY_PASS")

	fmt.Printf("user: %s, password: %s\n", user, password)
	db := store.NewFileStore("file", "products.json")
	//rep := products.NewRepository()
	rep := products.NewStoreRepository(db)
	service := products.NewService(rep)
	productHandler := handler.NewProduct(service)

	router := gin.Default()
	router.Use(TokenMiddleware)

	pr := router.Group("/products")
	pr.GET("/", productHandler.GetAll())       //OK
	pr.POST("/", productHandler.Create())      //OK
	pr.PUT("/:id", productHandler.Update())    //OK
	pr.PATCH("/", productHandler.UpdateFull()) //OK
	pr.DELETE("/:id", productHandler.Delete()) //OK
	router.Run()
}
