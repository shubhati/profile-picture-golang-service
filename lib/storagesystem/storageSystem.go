package storagesystem

type StorageSystem interface {
	DownloadFile(bucketName string, objectKey string) ([]byte, error)
}
