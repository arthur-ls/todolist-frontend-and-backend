package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load("resources/dbParams.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//port, _ := strconv.Atoi(os.Getenv("PG_PORT"))

	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PWD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))

	db, err := gorm.Open(mysql.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, err
}
