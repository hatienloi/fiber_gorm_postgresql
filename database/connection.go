package database

import (
	"github.com/hatienl0i261299/fiber_gorm_postgresql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("can not connect to database")
	}

	err = connection.AutoMigrate(
		&models.User{},
		&models.Article{},
	)

	if err != nil {
		panic("can not migrate database")
	}

	DB = connection
}
