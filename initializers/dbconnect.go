package initializers

import (
	"cse/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DATABASE_URL")
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	err = DB.AutoMigrate(
		&models.Project{},
		&models.File{},
		&models.Class{},
		&models.Method{},
		&models.Variable{},
		&models.Relationship{},
	)
	if err != nil {
		log.Fatalf("Unable to migrate the database schema: %v\n", err)
	}

	log.Println("Connected to the database and migrated schema successfully")
}
