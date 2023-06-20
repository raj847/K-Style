package utils

import (
	"kstyle-test/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	// Connect using gorm MySQL
	dsn := os.Getenv("DATABASE_URL") // Assuming the environment variable contains the MySQL connection string

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn.AutoMigrate(
		&entity.Member{},
		&entity.Product{},
		&entity.Like{},
		&entity.Review{},
	)

	SetupDBConnection(conn)
	return conn, nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
