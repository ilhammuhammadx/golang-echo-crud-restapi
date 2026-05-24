package config

import (
	"log"

	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	user := GetEnv("DB_USER", "postgres")
	password := GetEnv("DB_PASSWORD", "password")
	dbname := GetEnv("DB_NAME", "myapp")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}
