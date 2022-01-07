package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	Env         string
)

func SetEnv(environment string) {
	if environment == "local" {
		LoadEnv() //loading env
	}
	Env = environment
	log.Printf("Current env: %v", Env)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

func GetCORSConfig() cors.Config {
	cfg := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}
	return cfg
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//NewDatabase : intializes and returns mysql db
func NewMySQL() {
	var URL string
	if Env == "local" {
		URL = fmt.Sprintf("%v:%v@tcp(%v)/%v", os.Getenv("TEST_DB_USERNAME"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_URL"), os.Getenv("TEST_DB_NAME"))
	} else {
		URL = fmt.Sprintf("%v:%v@tcp(%v)/%v", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	}
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
	var (
		redisAddress  string
		redisPassword string
	)
	if Env == "local" {
		redisAddress = fmt.Sprintf("%v:%v", os.Getenv("TEST_REDIS_URL"), os.Getenv("TEST_REDIS_PORT"))
		redisPassword = os.Getenv("TEST_REDIS_PASSWORD")
	} else {
		redisAddress = fmt.Sprintf("%v:%v", os.Getenv("REDIS_URL"), os.Getenv("REDIS_PORT"))
		redisPassword = os.Getenv("REDIS_PASSWORD")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Printf("Error while establishing Redis Client: %v", err)
	}
	log.Println("NewRedisClient: Redis connection established")
	RedisClient = rdb
}
