package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// PageData represents data to be passed to templates
type PageData struct {
    Title    string
    Content  string
    Items    []Item
    Time     time.Time
    User     User
}

type Item struct {
    Name  string
    Price float64
}

type User struct {
    Name  string
    Email string
}

// Template functions
var funcMap = template.FuncMap{
    "formatTime": func(t time.Time) string {
        return t.Format("2006-01-02 15:04:05")
    },
    "formatPrice": func(price float64) string {
        return fmt.Sprintf("$%.2f", price)
    },
}

// Base template
const baseTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .container { max-width: 800px; margin: 0 auto; }
        .header { background: #f0f0f0; padding: 10px; }
        .content { margin-top: 20px; }
        .footer { margin-top: 20px; padding-top: 20px; border-top: 1px solid #ccc; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>{{.Title}}</h1>
        </div>
        <div class="content">
            {{template "content" .}}
        </div>
        <div class="footer">
            <p>Generated at {{.Time | formatTime}}</p>
        </div>
    </div>
</body>
</html>
`

// Content templates
const homeTemplate = `
{{define "content"}}
    <p>{{.Content}}</p>
    <h2>Featured Items</h2>
    <ul>
    {{range .Items}}
        <li>{{.Name}} - {{.Price | formatPrice}}</li>
    {{end}}
    </ul>
{{end}}
`

const userTemplate = `
{{define "content"}}
    <h2>User Profile</h2>
    <p><strong>Name:</strong> {{.User.Name}}</p>
    <p><strong>Email:</strong> {{.User.Email}}</p>
{{end}}
`

// Template caching
var templates = make(map[string]*template.Template)

func init() {
    // Parse and cache templates
    homeTmpl := template.New("home").Funcs(funcMap)
    homeTmpl = template.Must(homeTmpl.Parse(baseTemplate))
    homeTmpl = template.Must(homeTmpl.Parse(homeTemplate))
    templates["home"] = homeTmpl

    userTmpl := template.New("user").Funcs(funcMap)
    userTmpl = template.Must(userTmpl.Parse(baseTemplate))
    userTmpl = template.Must(userTmpl.Parse(userTemplate))
    templates["user"] = userTmpl
}

// Handler functions
func handleHome(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title:   "Welcome",
        Content: "Welcome to our website!",
        Items: []Item{
            {Name: "Item 1", Price: 19.99},
            {Name: "Item 2", Price: 29.99},
            {Name: "Item 3", Price: 39.99},
        },
        Time: time.Now(),
    }

    if err := templates["home"].Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handleUser(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title: "User Profile",
        User: User{
            Name:  "John Doe",
            Email: "john@example.com",
        },
        Time: time.Now(),
    }

    if err := templates["user"].Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Dynamic template example
func handleDynamic(w http.ResponseWriter, r *http.Request) {
    // Create a new template dynamically
    tmpl := template.New("dynamic").Funcs(funcMap)
    tmpl = template.Must(tmpl.Parse(`
        <!DOCTYPE html>
        <html>
        <head>
            <title>{{.Title}}</title>
        </head>
        <body>
            <h1>{{.Title}}</h1>
            <p>{{.Content}}</p>
            <p>Current time: {{.Time | formatTime}}</p>
        </body>
        </html>
    `))

    data := PageData{
        Title:   "Dynamic Page",
        Content: "This page was generated dynamically!",
        Time:    time.Now(),
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Template inheritance example
func handleInheritance(w http.ResponseWriter, r *http.Request) {
    // Create a new template with inheritance
    tmpl := template.New("inheritance").Funcs(funcMap)
    tmpl = template.Must(tmpl.Parse(baseTemplate))
    tmpl = template.Must(tmpl.Parse(`
        {{define "content"}}
            <h2>Inherited Content</h2>
            <p>This content inherits from the base template.</p>
            <p>{{.Content}}</p>
        {{end}}
    `))

    data := PageData{
        Title:   "Template Inheritance",
        Content: "This is inherited content!",
        Time:    time.Now(),
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    // Register routes
    http.HandleFunc("/", handleHome)
    http.HandleFunc("/user", handleUser)
    http.HandleFunc("/dynamic", handleDynamic)
    http.HandleFunc("/inheritance", handleInheritance)

    // Start server
    fmt.Println("Template server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
} 