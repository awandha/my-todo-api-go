package database

import (
	"fmt"
	"log"
	"os"

	"my-todo-api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func ConnectDatabase() {
	// Try to load .env locally; ignore error in cloud
	_ = godotenv.Load()

	// Prefer Railway MySQL vars; fall back to your DB_* vars
	host := getenv("MYSQLHOST", os.Getenv("DB_HOST"))
	port := getenv("MYSQLPORT", os.Getenv("DB_PORT"))
	user := getenv("MYSQLUSER", os.Getenv("DB_USER"))
	pass := getenv("MYSQLPASSWORD", os.Getenv("DB_PASSWORD"))
	name := getenv("MYSQLDATABASE", os.Getenv("DB_NAME"))

	if host == "" || user == "" || name == "" {
		log.Fatal("Database env vars are not set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Migrate all needed tables
	if err := db.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	DB = db
	log.Println("Connected to MySQL")
}
