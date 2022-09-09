package storagesystem

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Storage struct {
	s3Session *session.Session
}

func NewS3StorageSystem(region string) *S3Storage {
	sessionObj, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	return &S3Storage{
		s3Session: sessionObj,
	}
}

func (s *S3Storage) DownloadFile(bucketName string, objectKey string) ([]byte, error) {
	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}
	buf := aws.NewWriteAtBuffer([]byte{})

	downloader := s3manager.NewDownloader(awsSession)
	if _, err := downloader.Download(buf, requestInput); err != nil {
		log.Println("Failed to download s3 file", err.Error())
		return []byte{}, err
	}
	return buf.Bytes(), nil
}
