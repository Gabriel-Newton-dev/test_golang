package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {

	dbName := viper.Get("DB_NAME")
	dbPassword := viper.Get("DB_PASSWORD")
	dbUser := viper.Get("DB_USER")
	connection := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/ %s?charset=utf8mb4&parseTime=True", dbUser, dbPassword, dbName)
	_, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao se conectar no banco de dados", err)
	}
}

func GetDatabase() *gorm.DB {
	return DB
}
