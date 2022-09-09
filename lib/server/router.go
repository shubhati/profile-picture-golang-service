package server

import (
	"example.com/lib/controllers"
	"example.com/lib/storagesystem"
	"github.com/gin-gonic/gin"
)

/*
	new router returns the gin.Engine object after adding all the routes in it. At present we only have profile picture
	get route.
*/
func NewRouter() *gin.Engine {
	router := gin.New()
	addProfilePicRoutes(router)
	return router
}

/*
	addProfilePicRoutes adds the profile picture routes to the gin Engine
*/
func addProfilePicRoutes(router *gin.Engine) {
	localFileSystem := storagesystem.NewLocalStorageSystem()
	profilePicController := controllers.NewProfilePictureController(localFileSystem, "test-bucket")
	router.GET("/v1/profile-pic", profilePicController.GetProfilePic)
}
