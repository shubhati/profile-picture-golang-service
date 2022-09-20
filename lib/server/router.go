package server

import (
	"example.com/lib/controllers"
	"example.com/lib/storagesystem"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

/*
	new router returns the gin.Engine object after adding all the routes in it. At present we only have profile picture
	get route.
*/
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	addProfilePicRoutes(router)
	return router
}

/*
	addProfilePicRoutes adds the profile picture routes to the gin Engine
*/
func addProfilePicRoutes(router *mux.Router) {
	storageSystem := storagesystem.NewStorageSystem()
	profilePicController := controllers.NewProfilePictureController(storageSystem, viper.GetString("bucket"))

	profilePictureRoutes := router.PathPrefix("/v1/profile-pic").Subrouter()

	profilePictureRoutes.HandleFunc("", profilePicController.GetProfilePic).Methods("GET")
	profilePictureRoutes.HandleFunc("", profilePicController.UpdateProfilePicture).Methods("POST")
	profilePictureRoutes.HandleFunc("", profilePicController.DeleteProfilePicture).Methods("DELETE")
}
