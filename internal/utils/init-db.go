package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func InitDB() (*gorm.DB, error) {
	// dsn := "<username>:<password>@tcp(<hostname>:<port>)/<database>?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("MY_SQL_URI")

	// for now using sqlLite db
	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite", // Force GORM to use modernc SQLite
		DSN:        "app.db", // Database file
	}, &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
