package postgresql

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("POSTGRESQL_DB_HOST"), os.Getenv("POSTGRESQL_DB_USERNAME"), os.Getenv("POSTGRESQL_DB_PASSWORD"), os.Getenv("POSTGRESQL_DB_NAME"), os.Getenv("POSTGRESQL_DB_PORT"))
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
