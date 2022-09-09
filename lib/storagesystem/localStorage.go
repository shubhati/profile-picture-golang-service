package storagesystem

import (
	"fmt"
	"log"
	"os"
)

type LocalStorage struct {
}

func (s *LocalStorage) DownloadFile(bucketName string, objectKey string) ([]byte, error) {
	filename := fmt.Sprintf("./%s/%s", bucketName, objectKey)
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read file", err.Error())
		return []byte{}, err
	}
	return fileContent, nil
}

func NewLocalStorageSystem() *LocalStorage {
	return &LocalStorage{}
}
