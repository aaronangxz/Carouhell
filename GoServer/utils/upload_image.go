package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aaronangxz/TIC2601/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadBase64Image(item_id uint32, base64File string) (string, error) {
	var (
		s3Bucket = os.Getenv("S3_BUCKET")
		s3Region = os.Getenv("AWS_S3_REGION")
	)

	decode, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		return "", err
	}

	_, s3err := s3.New(models.S3Client).PutObject(&s3.PutObjectInput{
		Bucket:          aws.String(s3Bucket),
		Key:             aws.String("listing_" + fmt.Sprint(item_id) + ".jpg"),
		Body:            bytes.NewReader(decode),
		ContentEncoding: aws.String("base64"),
		ContentType:     aws.String("image/jpg"),
		ACL:             aws.String("public-read"),
	})

	if s3err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("https://%v.s3.%v.amazonaws.com/listing_%v.jpg", s3Bucket, s3Region, item_id)
	return fileName, nil
}
