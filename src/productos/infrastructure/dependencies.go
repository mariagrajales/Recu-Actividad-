package infrastructure

import (
	"233338-R-C2/src/productos/application"
	"github.com/gin-gonic/gin"
)

func ConfigureProductRoutes(r *gin.Engine) {

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	mysql := NewMySQL()

	// Casos de uso
	addProduct := application.NewAddProduct(mysql)
	getLastProduct := application.NewGetLastProduct(mysql)
	countProductsInDiscount := application.NewCountProductsInDiscount(mysql)
	listProduct := application.NewListProduct(mysql)

	// Controladores
	addProductController := NewAddProductController(addProduct)
	isNewProductAddedController := NewIsNewProductAddedController(getLastProduct)
	countProductsInDiscountController := NewCountProductsInDiscountController(countProductsInDiscount)
	listProductController := NewListProductController(listProduct)

	// Rutas
	api := r.Group("/api")
	{
		api.POST("/addProducto", addProductController.Execute)
		api.GET("/isNewProductAdded", isNewProductAddedController.Execute)
		api.GET("/countProductsInDiscount", countProductsInDiscountController.Execute)
		api.GET("/productos", listProductController.Execute)
	}
}
