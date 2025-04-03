package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func initDB() error {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    ctx := context.Background()
    return rdb.Ping(ctx).Err()
}

// String operations
func stringOperations(ctx context.Context) error {
    // Set value
    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        return err
    }

    // Get value
    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        return err
    }
    fmt.Printf("String value: %s\n", val)

    // Set with expiration
    err = rdb.Set(ctx, "expiring_key", "value", time.Hour).Err()
    if err != nil {
        return err
    }

    return nil
}

// List operations
func listOperations(ctx context.Context) error {
    // Push values
    err := rdb.LPush(ctx, "list", "world").Err()
    if err != nil {
        return err
    }
    err = rdb.LPush(ctx, "list", "hello").Err()
    if err != nil {
        return err
    }

    // Get all values
    values, err := rdb.LRange(ctx, "list", 0, -1).Result()
    if err != nil {
        return err
    }
    fmt.Printf("List values: %v\n", values)

    return nil
}

// Hash operations
func hashOperations(ctx context.Context) error {
    // Set hash fields
    err := rdb.HSet(ctx, "user:1", map[string]interface{}{
        "name":  "John",
        "email": "john@example.com",
        "age":   "30",
    }).Err()
    if err != nil {
        return err
    }

    // Get all fields
    all, err := rdb.HGetAll(ctx, "user:1").Result()
    if err != nil {
        return err
    }
    fmt.Printf("Hash values: %v\n", all)

    // Get specific field
    name, err := rdb.HGet(ctx, "user:1", "name").Result()
    if err != nil {
        return err
    }
    fmt.Printf("User name: %s\n", name)

    return nil
}

// Set operations
func setOperations(ctx context.Context) error {
    // Add members
    err := rdb.SAdd(ctx, "set", "member1", "member2", "member3").Err()
    if err != nil {
        return err
    }

    // Get all members
    members, err := rdb.SMembers(ctx, "set").Result()
    if err != nil {
        return err
    }
    fmt.Printf("Set members: %v\n", members)

    // Check membership
    exists, err := rdb.SIsMember(ctx, "set", "member1").Result()
    if err != nil {
        return err
    }
    fmt.Printf("member1 exists: %v\n", exists)

    return nil
}

// Sorted Set operations
func sortedSetOperations(ctx context.Context) error {
    // Add members with scores
    err := rdb.ZAdd(ctx, "scores", redis.Z{
        Score:  100.0,
        Member: "player1",
    }).Err()
    if err != nil {
        return err
    }
    err = rdb.ZAdd(ctx, "scores", redis.Z{
        Score:  200.0,
        Member: "player2",
    }).Err()
    if err != nil {
        return err
    }

    // Get members by score range
    members, err := rdb.ZRangeByScore(ctx, "scores", &redis.ZRangeBy{
        Min: "0",
        Max: "+inf",
    }).Result()
    if err != nil {
        return err
    }
    fmt.Printf("Sorted set members: %v\n", members)

    return nil
}

// Pub/Sub operations
func pubSubOperations(ctx context.Context) error {
    pubsub := rdb.Subscribe(ctx, "channel")
    defer pubsub.Close()

    // Start goroutine to receive messages
    go func() {
        for {
            msg, err := pubsub.ReceiveMessage(ctx)
            if err != nil {
                log.Printf("Error receiving message: %v", err)
                return
            }
            fmt.Printf("Received message: %s\n", msg.Payload)
        }
    }()

    // Publish message
    err := rdb.Publish(ctx, "channel", "hello").Err()
    if err != nil {
        return err
    }

    // Wait for message to be received
    time.Sleep(time.Second)
    return nil
}

func main() {
    if err := initDB(); err != nil {
        log.Fatal(err)
    }
    defer rdb.Close()

    ctx := context.Background()

    // Run examples
		if err := stringOperations(ctx); err != nil {
        log.Fatal(err)
    }

    if err := listOperations(ctx); err != nil {
        log.Fatal(err)
    }

    if err := hashOperations(ctx); err != nil {
        log.Fatal(err)
    }

    if err := setOperations(ctx); err != nil {
        log.Fatal(err)
    }

    if err := sortedSetOperations(ctx); err != nil {
        log.Fatal(err)
    }

    if err := pubSubOperations(ctx); err != nil {
        log.Fatal(err)
    }
} 