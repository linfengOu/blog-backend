package persistence

import (
	"fmt"
	"github/linfengOu/write-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// TODO flawed singleton
var db *gorm.DB

func InitDB() {
	if db != nil {
		log.Println("DB initialization is already done")
		return
	}
	user := config.Get(config.PostgresUser)
	pwd := config.Get(config.PostgresPassword)
	dbName := config.Get(config.PostgresDBname)
	port := config.Get(config.PostgresPort)
	sslMode := config.Get(config.PostgresSslmode)
	tz := config.Get(config.PostgresTimezone)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", user, pwd, dbName, port, sslMode, tz)
	dbInner, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("DB initialization failed")
	}
	// TODO connection pool here?
	db = dbInner
}

func GetDB() *gorm.DB {
	return db
}
