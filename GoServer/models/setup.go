package models

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB          *gorm.DB
	S3Client    *session.Session
	RedisClient *redis.Client
	Ctx         = context.TODO()
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-16881.c74.us-east-1-4.ec2.cloud.redislabs.com:16881",
		Password: "wr2PToMVbwaGmo1gd0b9ky4hlNbnwpZz", // no password set
		DB:       0,                                  // use default DB
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Printf("Error while establishing Redis Client: %v", err)
	}
	log.Println("NewRedisClient: Redis connection established")
	RedisClient = rdb
}
