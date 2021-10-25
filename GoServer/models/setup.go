package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB       *gorm.DB
	S3Client *session.Session
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

//NewDatabase : intializes and returns mysql db
func NewMySQL() {

	URL := "b0bc6fadb8432d:f25c7f6b@tcp(us-cdbr-east-04.cleardb.com:3306)/heroku_bdc39d4687a85d4"
	fmt.Println(URL)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		log.Println(err)
		panic("Failed to connect to database!")
	}

	log.Println("Database connection established")
	DB = db
}

func NewAWSInstance() {
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("AWS_SECRET_KEY")
	s3Region := os.Getenv("AWS_S3_REGION")

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")

	_, err := creds.Get()

	if err != nil {
		log.Println(err)
	}

	cfg := aws.NewConfig().WithRegion(s3Region).WithCredentials(creds)

	s3Connection, err := session.NewSession(cfg)

	if err != nil {
		log.Println(err)
	}
	S3Client = s3Connection
}
