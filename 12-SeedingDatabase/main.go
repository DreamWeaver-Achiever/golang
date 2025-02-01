package main

import (
        "log"

        "gorm.io/driver/postgres"
        "gorm.io/gorm"
)

// Define your models (e.g., User)
type User struct {
        gorm.Model
        Name  string
        Email string
}

func main() {
        // Connect to the database
        dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable" 
        db, err := gorm.Open(postgres.Open(dsn))
        if err != nil {
                log.Fatal("Failed to connect to database: ", err)
        }

        // Migrate the models to create the tables
        if err := db.AutoMigrate(&User{}); err != nil {
                log.Fatal("Failed to migrate models: ", err)
        }

        // Seed the database with initial data
        users := []User{
                {Name: "Alice", Email: "alice@example.com"},
                {Name: "Bob", Email: "bob@example.com"},
        }

        // Create the users in the database
        if err := db.Create(&users).Error; err != nil {
                log.Fatal("Failed to seed database: ", err)
        }

        log.Println("Database seeded successfully!")
}