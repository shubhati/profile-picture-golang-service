package storagesystem

import (
	"io"

	"github.com/spf13/viper"
)

/*
	This interface contains the function used to download file from filesystem. Currently two filesystems
	implement this interface, i.e. localfilesystem and aws s3 object storage. profile picture controller
	takes the object implementing this interface to download the files.
*/
type StorageSystem interface {
	DownloadFile(bucketName string, objectKey string) ([]byte, error)
	UploadFile(file io.Reader, bucketName string, objectKey string) error
	DeleteFile(bucketName string, objectKey string) error
}

const S3StorageType = "s3"
const MockedStorageType = "mocked"

func NewStorageSystem() StorageSystem {
	storageType := viper.GetString("storagesystem")
	if storageType == S3StorageType {
		return NewS3StorageSystem()
	} else if storageType == MockedStorageType {
		return NewLocalStorageSystem()
	} else {
		panic("unsupported storage type")
	}
}
