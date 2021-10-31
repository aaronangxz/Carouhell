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
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB       *gorm.DB
	S3Client *session.Session
	Redis    redis.Conn
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

//NewDatabase : intializes and returns mysql db
func NewMySQL() {

	URL := fmt.Sprintf("%v:%v@tcp(%v)/%v", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	log.Printf("Connecting to %v", URL)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		log.Printf("Error while establishing DB Connection: %v", err)
		panic("Failed to connect to database!")
	}

	log.Println("NewMySQL: Database connection established")
	DB = db
}

func NewAWSInstance() {
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("AWS_SECRET_KEY")
	s3Region := os.Getenv("AWS_S3_REGION")

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")

	_, err := creds.Get()

	if err != nil {
		log.Printf("Error while establishing S3 Credentials: %v", err)
	}

	cfg := aws.NewConfig().WithRegion(s3Region).WithCredentials(creds)

	s3Connection, err := session.NewSession(cfg)

	if err != nil {
		log.Printf("Error while establishing S3 Session: %v", err)
	}
	log.Println("NewAWSInstance: S3 connection established")
	S3Client = s3Connection
}

func NewRedis() {
	c, err := redis.DialURL(os.Getenv("REDIS_URL"), redis.DialTLSSkipVerify(true))
	if err != nil {
		log.Printf("Error while establishing Redis: %v", err)
	}
	Redis = c
}
