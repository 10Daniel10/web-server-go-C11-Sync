package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// Product representa la estructura de un producto
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var (
	products  []Product
	idCounter int
	mutex     sync.Mutex
)

func main() {
	// Agregar productos de prueba
	products = []Product{
		{
			ID:          1,
			Name:        "Cheese - St. Andre",
			Quantity:    60,
			CodeValue:   "S73191A",
			IsPublished: true,
			Expiration:  "12/04/2022",
			Price:       50.15,
		},
		{
			ID:          2,
			Name:        "Apples",
			Quantity:    100,
			CodeValue:   "A12345",
			IsPublished: true,
			Expiration:  "25/12/2022",
			Price:       1.99,
		},
	}

	router := gin.Default()

	// Ruta para agregar un producto (POST)
	router.POST("/products", addProduct)

	// Ruta para obtener un producto por ID (GET)
	router.GET("/products/:id", getProductByID)

	// Ruta para modificar un producto por ID (PATCH)
	router.PATCH("/products/:id", updateProduct)

	// Ruta para eliminar un producto por ID (DELETE)
	router.DELETE("/products/:id", deleteProduct)

	router.Run(":8080")
}

func getProductByID(c *gin.Context) {
	// Obtener el ID del producto de los parámetros de la URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	// Buscar el producto en la lista
	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	// Si no se encuentra el producto, responder con un error
	c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
}

func addProduct(c *gin.Context) {
	var newProduct Product

	// Decodificar el cuerpo JSON del request en una estructura Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el producto"})
		return
	}

	// Validar los campos del producto (como se hizo en la implementación anterior)

	// Generar un nuevo ID (como se hizo en la implementación anterior)

	// Agregar el nuevo producto a la lista (como se hizo en la implementación anterior)
}

func updateProduct(c *gin.Context) {
	// Obtener el ID del producto de los parámetros de la URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	// Buscar el producto en la lista
	for i, p := range products {
		if p.ID == id {
			// Decodificar el cuerpo JSON del request en una estructura Product
			var updatedProduct Product
			if err := c.ShouldBindJSON(&updatedProduct); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el producto"})
				return
			}

			// Actualizar solo los campos que se proporcionaron
			if updatedProduct.Name != "" {
				products[i].Name = updatedProduct.Name
			}
			if updatedProduct.Quantity > 0 {
				products[i].Quantity = updatedProduct.Quantity
			}
			if updatedProduct.CodeValue != "" {
				products[i].CodeValue = updatedProduct.CodeValue
			}
			// Actualizar otros campos según sea necesario

			c.JSON(http.StatusOK, products[i])
			return
		}
	}

	// Si no se encuentra el producto, responder con un error
	c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
}

func deleteProduct(c *gin.Context) {
	// Obtener el ID del producto de los parámetros de la URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
		return
	}

	// Buscar el producto en la lista y eliminarlo
	for i, p := range products {
		if p.ID == id {
			// Eliminar el producto de la lista
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
			return
		}
	}

	// Si no se encuentra el producto, responder con un error
	c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
}
