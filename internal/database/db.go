package database

import (
	"log"

	"github.com/Vostok1611/tasks-service/internal/task"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=secret dbname=postgres port=5400 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	//  Автоматически создаём таблицу tasks (если её нет)
	if err := db.AutoMigrate(&task.Task{}); err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}
