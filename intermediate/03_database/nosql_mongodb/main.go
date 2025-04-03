package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents a user document in MongoDB
type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string            `bson:"name"`
    Email     string            `bson:"email"`
    Age       int               `bson:"age"`
    Interests []string          `bson:"interests"`
    CreatedAt time.Time         `bson:"created_at"`
    UpdatedAt time.Time         `bson:"updated_at"`
}

// Post represents a blog post document
type Post struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    UserID    primitive.ObjectID `bson:"user_id"`
    Title     string            `bson:"title"`
    Content   string            `bson:"content"`
    Tags      []string          `bson:"tags"`
    CreatedAt time.Time         `bson:"created_at"`
    UpdatedAt time.Time         `bson:"updated_at"`
}

var client *mongo.Client
var db *mongo.Database

func initDB() error {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    
    // Connect to MongoDB
    var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return err
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return err
    }

    // Get database
    db = client.Database("blog")
    return nil
}

// Create a new user
func createUser(user *User) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    collection := db.Collection("users")
    result, err := collection.InsertOne(context.TODO(), user)
    if err != nil {
        return err
    }

    user.ID = result.InsertedID.(primitive.ObjectID)
    return nil
}

// Get user by ID
func getUser(id primitive.ObjectID) (*User, error) {
    collection := db.Collection("users")
    var user User
    
    err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}

// Create a new post
func createPost(post *Post) error {
    post.CreatedAt = time.Now()
    post.UpdatedAt = time.Now()

    collection := db.Collection("posts")
    result, err := collection.InsertOne(context.TODO(), post)
    if err != nil {
        return err
    }

    post.ID = result.InsertedID.(primitive.ObjectID)
    return nil
}

// Get posts by user ID with pagination
func getUserPosts(userID primitive.ObjectID, page, limit int64) ([]Post, error) {
    collection := db.Collection("posts")
    skip := (page - 1) * limit

    cursor, err := collection.Find(context.TODO(),
        bson.M{"user_id": userID},
        options.Find().SetSkip(skip).SetLimit(limit).SetSort(bson.D{{Key: "created_at", Value: -1}}),
    )
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    var posts []Post
    if err = cursor.All(context.TODO(), &posts); err != nil {
        return nil, err
    }

    return posts, nil
}

// Update user with atomic operation
func updateUser(id primitive.ObjectID, update bson.M) error {
    collection := db.Collection("users")
    
    update["updated_at"] = time.Now()
    _, err := collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": id},
        bson.M{"$set": update},
    )
    return err
}

// Delete user and their posts (transaction example)
func deleteUserAndPosts(id primitive.ObjectID) error {
    session, err := client.StartSession()
    if err != nil {
        return err
    }
    defer session.EndSession(context.TODO())

    _, err = session.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
        // Delete user
        if _, err := db.Collection("users").DeleteOne(ctx, bson.M{"_id": id}); err != nil {
            return nil, err
        }

        // Delete user's posts
        if _, err := db.Collection("posts").DeleteMany(ctx, bson.M{"user_id": id}); err != nil {
            return nil, err
        }

        return nil, nil
    })

    return err
}

func main() {
    if err := initDB(); err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.TODO())

    // Create a user
    user := &User{
        Name:      "John Doe",
        Email:     "john@example.com",
        Age:       30,
        Interests: []string{"golang", "mongodb", "web development"},
    }

    if err := createUser(user); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created user: %+v\n", user)

    // Create a post
    post := &Post{
        UserID:  user.ID,
        Title:   "Getting Started with MongoDB",
        Content: "MongoDB is a popular NoSQL database...",
        Tags:    []string{"mongodb", "nosql", "database"},
    }

    if err := createPost(post); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created post: %+v\n", post)

    // Get user's posts
    posts, err := getUserPosts(user.ID, 1, 10)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nUser's posts:")
    for _, p := range posts {
        fmt.Printf("- %s\n", p.Title)
    }

    // Update user
    update := bson.M{
        "age": 31,
        "interests": []string{"golang", "mongodb", "web development", "cloud"},
    }
    if err := updateUser(user.ID, update); err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nUpdated user")

    // Delete user and posts
    if err := deleteUserAndPosts(user.ID); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Deleted user and their posts")
} 