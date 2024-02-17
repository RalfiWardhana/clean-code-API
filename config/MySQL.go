package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/rumah_sakit?charset=utf8mb4&parseTime=True&loc=Local", CONFIG["MYSQL_USER"], CONFIG["MYSQL_PASS"], CONFIG["MYSQL_HOST"], CONFIG["MYSQL_PORT"])

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Println("connected")
	return db
}
