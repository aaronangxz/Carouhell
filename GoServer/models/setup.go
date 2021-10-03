package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Change CurrentENV : live / test
var (
	DB         *gorm.DB
	CurrentENV = "test"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

//NewDatabase : intializes and returns mysql db
func NewMySQL() {
	USER := os.Getenv("TEST_DB_USER")
	PASS := os.Getenv("TEST_DB_PASSWORD")
	//HOST := os.Getenv("TEST_DB_HOST")
	//PORT := os.Getenv("TEST_DB_PORT")
	DBNAME := os.Getenv("TEST_DB_NAME")

	//URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	URL := fmt.Sprintf("%s:%s@tcp(tic2601-db)/%s", USER, PASS, DBNAME)
	fmt.Println(URL)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection established")

	//db.AutoMigrate(&Listing{})
	//db.AutoMigrate(&Notification{})
	DB = db
}

func DBName() string {
	switch CurrentENV {
	case "test":
		return "tic2601_test_db"
	case "live":
		return "tic2601_db"
	default:
		return "tic2601_test_db"
	}
}

// func ConnectDataBase() {
// 	database, err := gorm.Open("sqlite3", "test.db")

// 	if err != nil {
// 		panic("Failed to connect to database!")
// 	}

// 	database.AutoMigrate(&Listing{})
// 	database.AutoMigrate(&Notification{})

// 	DB = database
// }
