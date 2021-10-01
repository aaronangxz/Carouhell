package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// //Database struct
// type Database struct {
// 	DB *gorm.DB
// }

// //NewDatabase : intializes and returns mysql db
// func NewMySQL() Database {
// 	USER := os.Getenv("TEST_DB_USER")
// 	PASS := os.Getenv("TEST_DB_PASSWORD")
// 	HOST := os.Getenv("TEST_DB_HOST")
// 	DBNAME := os.Getenv("TEST_DB_NAME")

// 	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
// 	fmt.Println(URL)
// 	db, err := gorm.Open(mysql.Open(URL))

// 	if err != nil {
// 		panic("Failed to connect to database!")

// 	}
// 	fmt.Println("Database connection established")
// 	return Database{
// 		DB: db,
// 	}

// }

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Listing{})
	database.AutoMigrate(&Notification{})

	DB = database
}
