package controllers

import (
	"log"
	"net/http"

	"example.com/lib/storagesystem"
	"github.com/gin-gonic/gin"
)

const IMAGE_TYPE = "image/jpeg"

/*
	ProfilePictureController is the rest controller which takes care of APIs related to the profile picture
	get query.
*/
type ProfilePictureController struct {
	storageSystem storagesystem.StorageSystem
	bucketName    string
}

/*
	returns new instance of profile pic controller. it takes bucketname and storage system object as arguement.
*/
func NewProfilePictureController(storageSystem storagesystem.StorageSystem, bucketName string) *ProfilePictureController {
	return &ProfilePictureController{
		storageSystem: storageSystem,
		bucketName:    bucketName,
	}
}

/*
	reads header named guid and finds the file named "guid.jpg" under the bucket name in the storage system.
*/
func (p *ProfilePictureController) GetProfilePic(c *gin.Context) {
	guid := c.GetHeader("guid")
	s3Key := guid + ".jpg"
	f, err := p.storageSystem.DownloadFile(p.bucketName, s3Key)
	if err != nil {
		log.Println("Failed to download the file:", err.Error())
		c.Error(err)
	}
	c.Data(http.StatusOK, IMAGE_TYPE, f)
}
