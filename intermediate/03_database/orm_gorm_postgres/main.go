package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Product represents a product in the database
type Product struct {
    ID          uint      `gorm:"primarykey"`
    Name        string    `gorm:"size:255;not null"`
    Price       float64   `gorm:"type:decimal(10,2);not null"`
    Description string    `gorm:"type:text"`
    Stock       int       `gorm:"default:0"`
    CategoryID  uint
    Category    Category  `gorm:"foreignKey:CategoryID"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"` // Soft delete
}

// Category represents a product category
type Category struct {
    ID        uint      `gorm:"primarykey"`
    Name      string    `gorm:"size:100;not null;unique"`
    Products  []Product
    CreatedAt time.Time
    UpdatedAt time.Time
}

// BeforeCreate is a GORM hook that runs before creating a product
func (p *Product) BeforeCreate(tx *gorm.DB) error {
    p.CreatedAt = time.Now()
    p.UpdatedAt = time.Now()
    return nil
}

// BeforeUpdate is a GORM hook that runs before updating a product
func (p *Product) BeforeUpdate(tx *gorm.DB) error {
    p.UpdatedAt = time.Now()
    return nil
}

var db *gorm.DB

func initDB() error {
    dsn := "host=localhost port=5432 user=postgres password=postgres dbname=products sslmode=disable"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // Enable SQL logging
    })
    if err != nil {
        return err
    }

    // Auto migrate the schema
    return db.AutoMigrate(&Category{}, &Product{})
}

func main() {
    if err := initDB(); err != nil {
        log.Fatal(err)
    }

    // Create a category
    category := Category{
        Name: "Electronics",
    }
    if err := db.Create(&category).Error; err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created category: %+v\n", category)

    // Create a product with category
    product := Product{
        Name:        "Smartphone",
        Price:       999.99,
        Description: "Latest smartphone model",
        Stock:       10,
        CategoryID:  category.ID,
    }
    if err := db.Create(&product).Error; err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created product: %+v\n", product)

    // Query with preload (eager loading)
    var products []Product
    if err := db.Preload("Category").Find(&products).Error; err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nAll products with categories:")
    for _, p := range products {
        fmt.Printf("- %s (Category: %s)\n", p.Name, p.Category.Name)
    }

    // Update product
    if err := db.Model(&product).Update("Price", 1099.99).Error; err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\nUpdated product price: %.2f\n", product.Price)

    // Query with conditions
    var expensiveProducts []Product
    if err := db.Where("price > ?", 1000).Find(&expensiveProducts).Error; err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nExpensive products:")
    for _, p := range expensiveProducts {
        fmt.Printf("- %s: $%.2f\n", p.Name, p.Price)
    }

    // Transaction example
    err := db.Transaction(func(tx *gorm.DB) error {
        // Update stock in transaction
        if err := tx.Model(&product).Update("Stock", gorm.Expr("Stock - ?", 1)).Error; err != nil {
            return err
        }
        // Create a new product in the same transaction
        newProduct := Product{
            Name:        "Tablet",
            Price:       499.99,
            Description: "New tablet model",
            Stock:       5,
            CategoryID:  category.ID,
        }
        return tx.Create(&newProduct).Error
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nTransaction completed successfully")

    // Soft delete
    if err := db.Delete(&product).Error; err != nil {
        log.Fatal(err)
    }
    fmt.Printf("\nSoft deleted product with ID: %d\n", product.ID)

    // Query including soft deleted records
    var allProducts []Product
    if err := db.Unscoped().Find(&allProducts).Error; err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nAll products (including soft deleted):")
    for _, p := range allProducts {
        fmt.Printf("- %s (Deleted: %v)\n", p.Name, p.DeletedAt.Valid)
    }
} 