package controllers

import (
	"log"
	"net/http"

	"example.com/lib/storagesystem"
	"github.com/gin-gonic/gin"
)

type ProfilePictureController struct {
	storageSystem storagesystem.StorageSystem
	bucketName    string
}

func NewProfilePictureController(storageSystem storagesystem.StorageSystem, bucketName string) *ProfilePictureController {
	return &ProfilePictureController{
		storageSystem: storageSystem,
		bucketName:    bucketName,
	}
}

func (p *ProfilePictureController) GetProfilePic(c *gin.Context) {
	guid := c.GetHeader("guid")
	s3Key := guid + ".jpg"
	f, err := p.storageSystem.DownloadFile(p.bucketName, s3Key)
	if err != nil {
		log.Println("Failed to download the file:", err.Error())
		c.Error(err)
	}
	c.Data(http.StatusOK, "image/jpeg", f)
}
