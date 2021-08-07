package utils

import (
	"log"
	"os"

	"github.com/Paulo-Lopes-Estevao/go_graphql_postegres/graph/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error     loading .env file %v", err)
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error  connect to datebase %v", err)
		panic(err)
	}

	//defer db.Close()

	db.AutoMigrate(&model.User)

	return db

}
