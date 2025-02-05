package adapters

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() {
	dsn := "host=localhost user=admin password=admin dbname=spark_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to Postgres")
		log.Fatal("Error connecting to Postgres:", err)
	}
	DB = db
	db.AutoMigrate(&Advice{})
}

type Advice struct {
	ID        int       `gorm:"primaryKey"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
}
