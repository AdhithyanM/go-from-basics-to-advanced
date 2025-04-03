package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// UserEvent represents an analytics event in Clickhouse
type UserEvent struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    EventType string    `gorm:"not null;type:String"`
    Path      string    `gorm:"type:String"`
    Device    string    `gorm:"type:String"`
    Country   string    `gorm:"type:String"`
    Timestamp time.Time `gorm:"not null"`
    Duration  float64   `gorm:"type:Float64"`
    // Clickhouse specific
    Date      time.Time `gorm:"type:Date;not null"`
}

// BeforeCreate hook to set the Date field from Timestamp
func (ue *UserEvent) BeforeCreate(tx *gorm.DB) error {
    ue.Date = ue.Timestamp.UTC().Truncate(24 * time.Hour)
    return nil
}

var db *gorm.DB

func initDB() error {
    dsn := "clickhouse://localhost:9000/analytics?dial_timeout=10s&read_timeout=20s"
    var err error
    
    db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    // Create the table with Clickhouse-specific options
    return db.Set("gorm:table_options", "ENGINE = MergeTree() PARTITION BY toYYYYMM(Date) ORDER BY (Timestamp, UserID) SETTINGS index_granularity = 8192").
        AutoMigrate(&UserEvent{})
}

// Insert batch of events
func insertBatchEvents(events []UserEvent) error {
    return db.CreateInBatches(events, 1000).Error
}

// Query events with aggregation
func getHourlyEventCounts(eventType string, startTime, endTime time.Time) ([]struct {
    Hour  time.Time
    Count int64
}, error) {
    var results []struct {
        Hour  time.Time
        Count int64
    }

    err := db.Table("user_events").
        Select("toStartOfHour(Timestamp) as Hour, count(*) as Count").
        Where("EventType = ? AND Timestamp BETWEEN ? AND ?", eventType, startTime, endTime).
        Group("Hour").
        Order("Hour").
        Find(&results).Error

    return results, err
}

// Get user session analytics
func getUserSessionAnalytics(startTime, endTime time.Time) ([]struct {
    Country       string
    AvgDuration   float64
    SessionCount  int64
}, error) {
    var results []struct {
        Country       string
        AvgDuration   float64
        SessionCount  int64
    }

    err := db.Table("user_events").
        Select("Country, avg(Duration) as AvgDuration, count(*) as SessionCount").
        Where("Timestamp BETWEEN ? AND ?", startTime, endTime).
        Group("Country").
        Having("SessionCount > ?", 100).
        Order("SessionCount DESC").
        Find(&results).Error

    return results, err
}

func main() {
    if err := initDB(); err != nil {
        log.Fatal(err)
    }

    // Generate sample events
    now := time.Now()
    events := []UserEvent{
        {
            UserID:    1,
            EventType: "pageview",
            Path:      "/home",
            Device:    "mobile",
            Country:   "US",
            Timestamp: now,
            Duration:  2.5,
        },
        {
            UserID:    2,
            EventType: "click",
            Path:      "/products",
            Device:    "desktop",
            Country:   "UK",
            Timestamp: now.Add(-1 * time.Hour),
            Duration:  1.8,
        },
        // Add more sample events...
    }

    // Insert events
    if err := insertBatchEvents(events); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Inserted sample events")

    // Query hourly event counts
    startTime := now.Add(-24 * time.Hour)
    endTime := now
    counts, err := getHourlyEventCounts("pageview", startTime, endTime)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nHourly pageview counts:")
    for _, count := range counts {
        fmt.Printf("%s: %d events\n", count.Hour.Format("15:04"), count.Count)
    }

    // Get session analytics
    analytics, err := getUserSessionAnalytics(startTime, endTime)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("\nUser session analytics by country:")
    for _, stat := range analytics {
        fmt.Printf("%s: %.2f avg duration, %d sessions\n",
            stat.Country, stat.AvgDuration, stat.SessionCount)
    }
} 