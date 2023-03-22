package database

import (
	"fmt"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/initializers"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		initializers.Config.DBHost,
		initializers.Config.DBUserName,
		initializers.Config.DBUserPassword,
		initializers.Config.DBName,
		initializers.Config.DBPort,
	)

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
