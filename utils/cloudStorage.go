package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var s3Con *s3.S3
var bucketName string

func ConnectToSpaces() error {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment variables")
	}

	key := os.Getenv("CLOUD_ACCESS")
	secret := os.Getenv("CLOUD_SECRET")
	url := os.Getenv("CLOUD_END_POINT")
	region := os.Getenv("CLOUD_REGION")
	bucketName = os.Getenv("CLOUD_BUCKET")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(url),
	})
	if err != nil {
		return err
	}

	s3Con = s3.New(sess)

	return nil
}

func UploadFileToSpaces(fileName string, fileHeader *multipart.FileHeader) error {
	// Open the file for use
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Printf("Cannot open file\n")
		return err
	}
	defer file.Close()

	// Read the file content into a buffer
	buffer, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Cannot read file\n")
		return err
	}

	// Config settings: this is where the upload magic happens
	_, err = s3Con.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucketName),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(len(buffer))),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}

func DeleteFileFromSpaces(keyName string) error {
	_, err := s3Con.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyName),
	})

	return err
}

func GenerateFileName(folderName, fileName string) (string, error) {
	currentTime := time.Now()

	unixTime := currentTime.Unix()

	unixTimeString := strconv.FormatInt(unixTime, 10)

	ext := filepath.Ext(fileName)

	fileNameWithoutExt := strings.TrimSuffix(fileName, ext)

	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	fileName = reg.ReplaceAllString(fileNameWithoutExt, "-")

	result := folderName + "/" + unixTimeString + "-" + fileName + ext

	return result, nil

}
