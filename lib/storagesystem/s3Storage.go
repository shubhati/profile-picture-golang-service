package storagesystem

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

/*
This structure implements StorageSystem interface and contains the functions to download the file from s3 bucket.
*/
type S3Storage struct {
	s3Session *session.Session
	s3Client  *s3.S3
}

/*
Creates the session on aws and returns a new object of type S3Storage.
*/
func NewS3StorageSystem() *S3Storage {
	region := viper.GetString("region")
	sessionObj, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	return &S3Storage{
		s3Session: sessionObj,
		s3Client:  s3.New(sessionObj),
	}
}

/*
	Downloads the file from s3 with the givne bucket name and object key. It returns the object in byte array type.
	if there's any error while downloading the file, the error is returned.
*/
func (s *S3Storage) DownloadFile(bucketName string, objectKey string) ([]byte, error) {
	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}
	buf := aws.NewWriteAtBuffer([]byte{})

	downloader := s3manager.NewDownloader(s.s3Session)
	if _, err := downloader.Download(buf, requestInput); err != nil {
		log.Println("Failed to download s3 file", err.Error())
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func (s *S3Storage) UploadFile(localFilePath string, bucketName string, objectKey string) error {
	file, err := os.Open(localFilePath)
	if err != nil {
		log.Println("Error in reading local file: ", err.Error())
		return err
	}
	defer file.Close()
	uploadInput := &s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	}
	uploader := s3manager.NewUploader(s.s3Session)
	uploaderOutput, err := uploader.Upload(uploadInput)
	if err != nil {
		log.Println("Error while uploading the file to s3: ", err.Error())
		return err
	}
	log.Println("File uploaded successfully, upload id:", uploaderOutput.UploadID)
	return nil
}

func (s *S3Storage) DeleteFile(bucketName string, objectKey string) error {
	deleteInput := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}
	result, err := s.s3Client.DeleteObject(deleteInput)
	if err != nil {
		log.Println("Error while deleting the object from s3:", err.Error())
		return err
	}
	log.Println("object deleted from s3, ", result.String())
	return nil
}
