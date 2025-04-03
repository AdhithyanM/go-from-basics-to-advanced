package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Product represents a product in the database
type Product struct {
    ID          int
    Name        string
    Price       float64
    Description string
    CreatedAt   time.Time
}

// Database connection
var db *sql.DB

// Initialize database
func initDB() error {
    var err error
    db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=products sslmode=disable")
    if err != nil {
        return err
    }

    // Test the connection
    if err := db.Ping(); err != nil {
        return err
    }

    // Create products table
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS products (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            price DECIMAL(10,2) NOT NULL,
            description TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    return err
}

// Create a new product using prepared statement
func createProduct(product *Product) error {
    query := `
        INSERT INTO products (name, price, description)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `
    return db.QueryRow(query, product.Name, product.Price, product.Description).
        Scan(&product.ID, &product.CreatedAt)
}

// Get a product by ID
func getProduct(id int) (*Product, error) {
    product := &Product{}
    query := `
        SELECT id, name, price, description, created_at
        FROM products
        WHERE id = $1
    `
    err := db.QueryRow(query, id).Scan(
        &product.ID,
        &product.Name,
        &product.Price,
        &product.Description,
        &product.CreatedAt,
    )
    if err != nil {
        return nil, err
    }
    return product, nil
}

// Get all products
func getProducts() ([]Product, error) {
    query := `
        SELECT id, name, price, description, created_at
        FROM products
    `
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var p Product
        err := rows.Scan(
            &p.ID,
            &p.Name,
            &p.Price,
            &p.Description,
            &p.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, rows.Err()
}

// Update a product
func updateProduct(product *Product) error {
    query := `
        UPDATE products
        SET name = $1, price = $2, description = $3
        WHERE id = $4
        RETURNING created_at
    `
    return db.QueryRow(query, product.Name, product.Price, product.Description, product.ID).
        Scan(&product.CreatedAt)
}

// Delete a product
func deleteProduct(id int) error {
    query := `DELETE FROM products WHERE id = $1`
    result, err := db.Exec(query, id)
    if err != nil {
        return err
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rows == 0 {
        return fmt.Errorf("product not found")
    }

    return nil
}

// Transaction example
func transferStock(fromID, toID, quantity int) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Deduct stock from source
    _, err = tx.Exec(`
        UPDATE products
        SET stock = stock - $1
        WHERE id = $2 AND stock >= $1
    `, quantity, fromID)
    if err != nil {
        return err
    }

    // Add stock to destination
    _, err = tx.Exec(`
        UPDATE products
        SET stock = stock + $1
        WHERE id = $2
    `, quantity, toID)
    if err != nil {
        return err
    }

    return tx.Commit()
}

// Connection pool example
func initConnectionPool() {
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)
}

func main() {
    // Initialize database
    if err := initDB(); err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize connection pool
    initConnectionPool()

    // Example usage
    product := &Product{
        Name:        "Example Product",
        Price:       29.99,
        Description: "This is an example product",
    }

    // Create product
    if err := createProduct(product); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created product: %+v\n", product)

    // Get product
    retrieved, err := getProduct(product.ID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Retrieved product: %+v\n", retrieved)

    // Update product
    product.Price = 39.99
    if err := updateProduct(product); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Updated product: %+v\n", product)

    // Get all products
    products, err := getProducts()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("All products:")
    for _, p := range products {
        fmt.Printf("- %+v\n", p)
    }

    // Delete product
    if err := deleteProduct(product.ID); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Deleted product with ID: %d\n", product.ID)
} 