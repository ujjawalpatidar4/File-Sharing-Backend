package config
import (
	"fmt"
	"log"
"file-sharing-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() {
	dsn := "host=localhost user=postgres password=Ujjawal@7613 dbname=file_sharing port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.File{})

	fmt.Println("Connected to Database!")
	DB = db
}

