package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// User represents a user item in DynamoDB
type User struct {
    UserID    string    `dynamodbav:"user_id"`
    Name      string    `dynamodbav:"name"`
    Email     string    `dynamodbav:"email"`
    Age       int       `dynamodbav:"age"`
    Interests []string  `dynamodbav:"interests"`
    CreatedAt time.Time `dynamodbav:"created_at"`
    UpdatedAt time.Time `dynamodbav:"updated_at"`
}

var client *dynamodb.Client

func initDB() error {
    // Load AWS configuration
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion("us-east-1"),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
            "your-access-key",
            "your-secret-key",
            "",
        )),
    )
    if err != nil {
        return err
    }

    // Create DynamoDB client
    client = dynamodb.NewFromConfig(cfg)
    return nil
}

// Create a new user
func createUser(user *User) error {
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    input := &dynamodb.PutItemInput{
        TableName: aws.String("users"),
        Item: map[string]types.AttributeValue{
            "user_id":    &types.AttributeValueMemberS{Value: user.UserID},
            "name":      &types.AttributeValueMemberS{Value: user.Name},
            "email":     &types.AttributeValueMemberS{Value: user.Email},
            "age":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", user.Age)},
            "interests": &types.AttributeValueMemberSS{Value: user.Interests},
            "created_at": &types.AttributeValueMemberS{Value: user.CreatedAt.Format(time.RFC3339)},
            "updated_at": &types.AttributeValueMemberS{Value: user.UpdatedAt.Format(time.RFC3339)},
        },
    }

    _, err := client.PutItem(context.TODO(), input)
    return err
}

// Get user by ID
func getUser(userID string) (*User, error) {
    input := &dynamodb.GetItemInput{
        TableName: aws.String("users"),
        Key: map[string]types.AttributeValue{
            "user_id": &types.AttributeValueMemberS{Value: userID},
        },
    }

    result, err := client.GetItem(context.TODO(), input)
    if err != nil {
        return nil, err
    }

    if result.Item == nil {
        return nil, fmt.Errorf("user not found")
    }

    user := &User{
        UserID: result.Item["user_id"].(*types.AttributeValueMemberS).Value,
        Name:   result.Item["name"].(*types.AttributeValueMemberS).Value,
        Email:  result.Item["email"].(*types.AttributeValueMemberS).Value,
        Age:    parseInt(result.Item["age"].(*types.AttributeValueMemberN).Value),
        Interests: result.Item["interests"].(*types.AttributeValueMemberSS).Value,
    }

    user.CreatedAt, _ = time.Parse(time.RFC3339, result.Item["created_at"].(*types.AttributeValueMemberS).Value)
    user.UpdatedAt, _ = time.Parse(time.RFC3339, result.Item["updated_at"].(*types.AttributeValueMemberS).Value)

    return user, nil
}

// Update user
func updateUser(userID string, updates map[string]types.AttributeValue) error {
    updates["updated_at"] = &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)}

    input := &dynamodb.UpdateItemInput{
        TableName: aws.String("users"),
        Key: map[string]types.AttributeValue{
            "user_id": &types.AttributeValueMemberS{Value: userID},
        },
        UpdateExpression:          aws.String("SET #name = :name, #email = :email, #age = :age, #interests = :interests, #updated_at = :updated_at"),
        ExpressionAttributeNames: map[string]string{
            "#name":       "name",
            "#email":      "email",
            "#age":        "age",
            "#interests":  "interests",
            "#updated_at": "updated_at",
        },
        ExpressionAttributeValues: updates,
    }

    _, err := client.UpdateItem(context.TODO(), input)
    return err
}

// Delete user
func deleteUser(userID string) error {
    input := &dynamodb.DeleteItemInput{
        TableName: aws.String("users"),
        Key: map[string]types.AttributeValue{
            "user_id": &types.AttributeValueMemberS{Value: userID},
        },
    }

    _, err := client.DeleteItem(context.TODO(), input)
    return err
}

// Query users by age range
func queryUsersByAge(minAge, maxAge int) ([]User, error) {
    input := &dynamodb.QueryInput{
        TableName:              aws.String("users"),
        IndexName:             aws.String("age-index"),
        KeyConditionExpression: aws.String("#age BETWEEN :min_age AND :max_age"),
        ExpressionAttributeNames: map[string]string{
            "#age": "age",
        },
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":min_age": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", minAge)},
            ":max_age": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", maxAge)},
        },
    }

    result, err := client.Query(context.TODO(), input)
    if err != nil {
        return nil, err
    }

    var users []User
    for _, item := range result.Items {
        user := &User{
            UserID: item["user_id"].(*types.AttributeValueMemberS).Value,
            Name:   item["name"].(*types.AttributeValueMemberS).Value,
            Email:  item["email"].(*types.AttributeValueMemberS).Value,
            Age:    parseInt(item["age"].(*types.AttributeValueMemberN).Value),
            Interests: item["interests"].(*types.AttributeValueMemberSS).Value,
        }
        user.CreatedAt, _ = time.Parse(time.RFC3339, item["created_at"].(*types.AttributeValueMemberS).Value)
        user.UpdatedAt, _ = time.Parse(time.RFC3339, item["updated_at"].(*types.AttributeValueMemberS).Value)
        users = append(users, *user)
    }

    return users, nil
}

func parseInt(s string) int {
    var i int
    fmt.Sscanf(s, "%d", &i)
    return i
}

func main() {
    if err := initDB(); err != nil {
        log.Fatal(err)
    }

    // Create a user
    user := &User{
        UserID:    "user1",
        Name:      "John Doe",
        Email:     "john@example.com",
        Age:       30,
        Interests: []string{"golang", "aws", "cloud"},
    }

    if err := createUser(user); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created user: %+v\n", user)

    // Get user
    retrievedUser, err := getUser(user.UserID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Retrieved user: %+v\n", retrievedUser)

    // Update user
    updates := map[string]types.AttributeValue{
        ":name":      &types.AttributeValueMemberS{Value: "John Smith"},
        ":email":     &types.AttributeValueMemberS{Value: "john.smith@example.com"},
        ":age":       &types.AttributeValueMemberN{Value: "31"},
        ":interests": &types.AttributeValueMemberSS{Value: []string{"golang", "aws", "cloud", "dynamodb"}},
    }

    if err := updateUser(user.UserID, updates); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Updated user")

    // Query users by age
    users, err := queryUsersByAge(25, 35)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nUsers aged 25-35:")
    for _, u := range users {
        fmt.Printf("- %s (age: %d)\n", u.Name, u.Age)
    }

    // Delete user
    if err := deleteUser(user.UserID); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Deleted user")
} 