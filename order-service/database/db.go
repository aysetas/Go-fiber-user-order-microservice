package database

import (
	"fmt"
	"log"

	"order-service/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	user := config.GetEnv("DB_USER")
	pass := config.GetEnv("DB_PASS")
	dbName := config.GetEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanırken hata oluştu:", err)
	}

	DB = db
	fmt.Println("✔ Veritabanı bağlantısı başarılı!")
}
