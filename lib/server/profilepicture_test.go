package server_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/lib/controllers"
	"example.com/lib/server"
	"github.com/gavv/httpexpect"
	"github.com/spf13/viper"
)

func TestProfilePictureServer(t *testing.T) {
	viper.Set("bucket", "../../test-bucket")
	router := server.NewRouter()
	testserver := httptest.NewServer(router)
	defer testserver.Close()

	testGuid := "test-guid"
	e := httpexpect.New(t, testserver.URL)

	log.Println("Creating a profile picture first")
	e.POST(server.PROFILE_PICTURE_PATH).WithHeader(controllers.GUID_HEADER, testGuid).
		WithMultipart().
		WithFile("photo", "../../test-bucket/example.jpg").
		Expect().
		Status(http.StatusOK)

	log.Println("recieving the created profile picture")

	e.GET(server.PROFILE_PICTURE_PATH).WithHeader(controllers.GUID_HEADER, testGuid).
		Expect().
		Status(http.StatusOK)

	log.Println("deleting the created profile picture")

	e.DELETE(server.PROFILE_PICTURE_PATH).WithHeader(controllers.GUID_HEADER, testGuid).
		Expect().
		Status(http.StatusOK)
}
