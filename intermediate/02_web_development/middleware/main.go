package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Call the next handler
        next(w, r)
        
        // Log the request details
        log.Printf(
            "%s %s %s %v",
            r.Method,
            r.RequestURI,
            r.RemoteAddr,
            time.Since(start),
        )
    }
}

// Authentication middleware
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Check for API key in header
        apiKey := r.Header.Get("X-API-Key")
        if apiKey != "secret123" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        next(w, r)
    }
}

// CORS middleware
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // Handle preflight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next(w, r)
    }
}

// Request ID middleware
func requestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Generate a request ID
        requestID := fmt.Sprintf("%d", time.Now().UnixNano())
        
        // Add request ID to response header
        w.Header().Set("X-Request-ID", requestID)
        
        // Add request ID to request context
        r.Header.Set("X-Request-ID", requestID)
        
        next(w, r)
    }
}

// Rate limiting middleware
func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
    // Simple in-memory rate limiter
    requests := make(map[string][]time.Time)
    
    return func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        now := time.Now()
        
        // Clean old requests
        if times, exists := requests[ip]; exists {
            var valid []time.Time
            for _, t := range times {
                if now.Sub(t) < time.Minute {
                    valid = append(valid, t)
                }
            }
            requests[ip] = valid
            
            // Check rate limit (10 requests per minute)
            if len(valid) >= 10 {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }
        }
        
        // Add current request
        requests[ip] = append(requests[ip], now)
        
        next(w, r)
    }
}

// Recovery middleware
func recoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Panic recovered: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        
        next(w, r)
    }
}

// Chain middleware
func chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for i := len(middlewares) - 1; i >= 0; i-- {
        f = middlewares[i](f)
    }
    return f
}

// Handler functions
func handleHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Home Page!")
}

func handleProtected(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is a protected endpoint!")
}

func handlePanic(w http.ResponseWriter, r *http.Request) {
    panic("This is a panic!")
}

func main() {
    // Create handlers with middleware
    homeHandler := chain(
        handleHome,
        loggingMiddleware,
        corsMiddleware,
        requestIDMiddleware,
        recoveryMiddleware,
    )
    
    protectedHandler := chain(
        handleProtected,
        loggingMiddleware,
        corsMiddleware,
        requestIDMiddleware,
        authMiddleware,
        recoveryMiddleware,
    )
    
    panicHandler := chain(
        handlePanic,
        loggingMiddleware,
        corsMiddleware,
        requestIDMiddleware,
        recoveryMiddleware,
    )
    
    // Register routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/protected", protectedHandler)
    http.HandleFunc("/panic", panicHandler)
    
    // Start server
    fmt.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
} 