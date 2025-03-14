package initializers

import (
	"fmt"
	"os"
	"simplechatagent-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Print(DB)

	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}

func SyncDB() {
	DB.AutoMigrate((&models.Post{}))
}
