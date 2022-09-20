package controllers

import (
	"fmt"
	"log"
	"net/http"

	"example.com/lib/storagesystem"
)

const IMAGE_TYPE = "image/jpeg"
const GUID_HEADER = "guid"
const PROFILE_PHOTO_KEY = "photo"

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
func (p *ProfilePictureController) GetProfilePic(rw http.ResponseWriter, r *http.Request) {
	guid := r.Header.Get(GUID_HEADER)
	s3Key := guid + ".jpg"
	f, err := p.storageSystem.DownloadFile(p.bucketName, s3Key)
	if err != nil {
		log.Println("Failed to download the file:", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, err.Error())
		return
	}
	rw.Header().Set("Content-Type", IMAGE_TYPE)
	if _, err = rw.Write(f); err != nil {
		log.Println("Error in writing response: " + err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, err.Error())
	}
}

func (p *ProfilePictureController) UpdateProfilePicture(rw http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile(PROFILE_PHOTO_KEY)
	if err != nil {
		log.Println("Error in reading file:", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, err.Error())
	}
	defer file.Close()

	guid := r.Header.Get(GUID_HEADER)
	s3Key := guid + ".jpg"
	err = p.storageSystem.UploadFile(file, p.bucketName, s3Key)
	if err != nil {
		log.Println("Error in uploading file: ", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, err.Error())
		return
	}
	// c.String(http.StatusOK, "")
}

func (p *ProfilePictureController) DeleteProfilePicture(rw http.ResponseWriter, r *http.Request) {
	guid := r.Header.Get(GUID_HEADER)
	s3Key := guid + ".jpg"
	err := p.storageSystem.DeleteFile(p.bucketName, s3Key)
	if err != nil {
		log.Println("Failed to delete the file:", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(rw, err.Error())
		return
	}
	// c.String(http.StatusOK, "")
}
