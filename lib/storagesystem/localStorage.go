package storagesystem

import (
	"fmt"
	"log"
	"os"
)

/*
This structure implements StorageSystem interface and contains the functions to download the file from local filesystem.
*/
type LocalStorage struct {
}

/*
	downloads the file from local directory with the name as "./bucketname/objectkey". It returns the object in byte array type.
	if there's any error while downloading the file, the error is returned.
*/
func (s *LocalStorage) DownloadFile(bucketName string, objectKey string) ([]byte, error) {
	filename := fmt.Sprintf("./%s/%s", bucketName, objectKey)
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read file", err.Error())
		return []byte{}, err
	}
	return fileContent, nil
}

/*
	returns new isntance of a LocalStorage structure.
*/
func NewLocalStorageSystem() *LocalStorage {
	return &LocalStorage{}
}
