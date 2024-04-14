package main

import (
	"log"
	"os"

	"github.com/davimerotto/web-server/cmd/server/handler"
	"github.com/davimerotto/web-server/cmd/server/middleware"
	"github.com/davimerotto/web-server/internal/products"
	"github.com/davimerotto/web-server/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/davimerotto/web-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	//user := os.Getenv("MY_USER")
	//password := os.Getenv("MY_PASS")

	//fmt.Printf("user: %s, password: %s\n", user, password)

	db := store.NewFileStore("file", "products.json")
	//rep := products.NewRepository()
	rep := products.NewStoreRepository(db)
	service := products.NewService(rep)
	productHandler := handler.NewProduct(service)

	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.TokenMiddleware)

	pr := router.Group("/products")
	pr.GET("/", productHandler.GetAll())       //OK
	pr.POST("/", productHandler.Create())      //OK
	pr.PUT("/:id", productHandler.Update())    //OK
	pr.PATCH("/", productHandler.UpdateFull()) //OK
	pr.DELETE("/:id", productHandler.Delete()) //OK
	router.Run()
}
