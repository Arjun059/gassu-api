package utils

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// dsn := "<username>:<password>@tcp(<hostname>:<port>)/<database>?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("MY_SQL_URI")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
