package storagesystem

import (
	"fmt"
	"io/ioutil"
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

func (s *LocalStorage) UploadFile(localFilePath string, bucketName string, objectKey string) error {
	file, err := os.Open(localFilePath)
	if err != nil {
		log.Println("Error in reading local file: ", err.Error())
		return err
	}
	defer file.Close()

	filename := fmt.Sprintf("./%s/%s", bucketName, objectKey)
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error in reading file: ", err.Error())
		return err
	}
	err = ioutil.WriteFile(filename, fileContent, 0755)
	if err != nil {
		log.Println("Error in saving the file locally: ", err.Error())
		return err
	}
	return nil
}

func (s *LocalStorage) DeleteFile(bucketName string, objectKey string) error {
	filename := fmt.Sprintf("./%s/%s", bucketName, objectKey)
	if err := os.Remove(filename); err != nil {
		log.Println("Error while removing the file: ", err.Error())
		return err
	}
	return nil
}

/*
	returns new isntance of a LocalStorage structure.
*/
func NewLocalStorageSystem() *LocalStorage {
	return &LocalStorage{}
}
