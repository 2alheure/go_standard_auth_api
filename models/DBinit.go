package models

import (
    "os"
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func DBInit() (error) {
    err := godotenv.Load("my.env")

    if err != nil {
        return err
    }

    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // TODO
    // Build DB URI in order to connect
    dbURI := fmt.Sprintf("%s:%s@%s:%S/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

    DB, err = gorm.Open("mysql", dbURI)
    
    return err
}