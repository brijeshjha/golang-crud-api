package main

import (
	"fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
	"net/http"
	"strconv"
	"os"
)

type Product struct {
    ID        int64 `json:"id"`
    ProductName string `json:"product_name"`
    Price      float64 `json:"price"`
    Description string `json:"description"`
}

var db *gorm.DB

func main() {

	host := os.Getenv("DB_HOST")
	fmt.Printf("host = %s", host)

    // Connect to PostgreSQL database
    dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		"5432",
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}


	db = database

    // Create a Gin router
    router := gin.Default()

    // Define routes for CRUD operations
    router.POST("/products", createProduct)
    router.GET("/products", getAllProducts)
    router.GET("/products/:id", getProductByID)
    router.PUT("/products/:id", updateProduct)
    router.DELETE("/products/:id", deleteProduct)

    // Start the Gin server
    router.Run(":3000")
}

func createProduct(c *gin.Context) {
    var product Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Insert product into database
    db.Create(&product)

    c.JSON(http.StatusOK, gin.H{"message": "Product created"})
}

func getAllProducts(c *gin.Context) {
    products := []Product{}

    if err := db.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    product := Product{}
    if err := db.Where("id = ?", id).First(&product).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, product)
}

func updateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    var product Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product.ID = int64(id)

    // Update product in database
    db.Save(&product)

    c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func deleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    // Delete product from database
    if err := db.Delete(&Product{}, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }
	}
}