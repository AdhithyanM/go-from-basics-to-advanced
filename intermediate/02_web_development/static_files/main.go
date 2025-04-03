package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// FileInfo represents information about a file
type FileInfo struct {
    Name    string
    Size    int64
    ModTime time.Time
    IsDir   bool
}

// FileCache represents a cache for static files
type FileCache struct {
    files map[string][]byte
    times map[string]time.Time
}

// Create a new file cache
func newFileCache() *FileCache {
    return &FileCache{
        files: make(map[string][]byte),
        times: make(map[string]time.Time),
    }
}

// Add a file to the cache
func (c *FileCache) add(path string, data []byte) {
    c.files[path] = data
    c.times[path] = time.Now()
}

// Get a file from the cache
func (c *FileCache) get(path string) ([]byte, bool) {
    data, exists := c.files[path]
    return data, exists
}

// Check if a file needs to be refreshed
func (c *FileCache) needsRefresh(path string, modTime time.Time) bool {
    cachedTime, exists := c.times[path]
    if !exists {
        return true
    }
    return modTime.After(cachedTime)
}

// Global file cache
var cache = newFileCache()

// Serve static file with caching
func handleStaticFile(w http.ResponseWriter, r *http.Request) {
    path := filepath.Join("static", r.URL.Path)
    
    // Check if file exists
    info, err := os.Stat(path)
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }
    
    // Check if file is in cache and needs refresh
    if data, exists := cache.get(path); exists && !cache.needsRefresh(path, info.ModTime()) {
        serveFile(w, path, data, info)
        return
    }
    
    // Read and cache file
    data, err := os.ReadFile(path)
    if err != nil {
        http.Error(w, "Error reading file", http.StatusInternalServerError)
        return
    }
    
    cache.add(path, data)
    serveFile(w, path, data, info)
}

// Serve file with appropriate headers
func serveFile(w http.ResponseWriter, path string, data []byte, info os.FileInfo) {
    // Set content type based on file extension
    switch filepath.Ext(path) {
    case ".css":
        w.Header().Set("Content-Type", "text/css")
    case ".js":
        w.Header().Set("Content-Type", "application/javascript")
    case ".jpg", ".jpeg":
        w.Header().Set("Content-Type", "image/jpeg")
    case ".png":
        w.Header().Set("Content-Type", "image/png")
    case ".gif":
        w.Header().Set("Content-Type", "image/gif")
    case ".svg":
        w.Header().Set("Content-Type", "image/svg+xml")
    }
    
    // Set caching headers
    w.Header().Set("Cache-Control", "public, max-age=3600")
    w.Header().Set("Last-Modified", info.ModTime().Format(http.TimeFormat))
    
    w.Write(data)
}

// Handle file upload
func handleUpload(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    // Parse multipart form
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }
    
    // Get file from form
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error getting file", http.StatusBadRequest)
        return
    }
    defer file.Close()
    
    // Create upload directory if it doesn't exist
    if err := os.MkdirAll("uploads", 0755); err != nil {
        http.Error(w, "Error creating directory", http.StatusInternalServerError)
        return
    }
    
    // Create destination file
    dst, err := os.Create(filepath.Join("uploads", handler.Filename))
    if err != nil {
        http.Error(w, "Error creating file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    
    // Copy file
    if _, err := io.Copy(dst, file); err != nil {
        http.Error(w, "Error copying file", http.StatusInternalServerError)
        return
    }
    
    fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

// Handle file download
func handleDownload(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Query().Get("file")
    if filename == "" {
        http.Error(w, "No file specified", http.StatusBadRequest)
        return
    }
    
    path := filepath.Join("uploads", filename)
    
    // Check if file exists
    info, err := os.Stat(path)
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }
    
    // Set headers for download
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
    w.Header().Set("Content-Type", "application/octet-stream")
    w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))
    
    // Serve file
    http.ServeFile(w, r, path)
}

// List uploaded files
func handleListFiles(w http.ResponseWriter, r *http.Request) {
    files, err := os.ReadDir("uploads")
    if err != nil {
        http.Error(w, "Error reading directory", http.StatusInternalServerError)
        return
    }
    
    var fileList []FileInfo
    for _, file := range files {
        info, err := file.Info()
        if err != nil {
            continue
        }
        fileList = append(fileList, FileInfo{
            Name:    info.Name(),
            Size:    info.Size(),
            ModTime: info.ModTime(),
            IsDir:   info.IsDir(),
        })
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(fileList)
}

func main() {
    // Create necessary directories
    os.MkdirAll("static", 0755)
    os.MkdirAll("uploads", 0755)
    
    // Register routes
    http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(handleStaticFile)))
    http.HandleFunc("/upload", handleUpload)
    http.HandleFunc("/download", handleDownload)
    http.HandleFunc("/files", handleListFiles)
    
    // Start server
    fmt.Println("Static file server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
} 