package storagesystem

/*
	This interface contains the function used to download file from filesystem. Currently two filesystems
	implement this interface, i.e. localfilesystem and aws s3 object storage. profile picture controller
	takes the object implementing this interface to download the files.
*/
type StorageSystem interface {
	DownloadFile(bucketName string, objectKey string) ([]byte, error)
}
