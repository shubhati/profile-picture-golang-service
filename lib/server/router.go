package server

import (
	"example.com/lib/controllers"
	"example.com/lib/storagesystem"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	addProfilePicRoutes(router)
	return router
}

func addProfilePicRoutes(router *gin.Engine) {
	localFileSystem := storagesystem.NewLocalStorageSystem()
	profilePicController := controllers.NewProfilePictureController(localFileSystem, "test-bucket")
	router.GET("/v1/profile-pic", profilePicController.GetProfilePic)
}
