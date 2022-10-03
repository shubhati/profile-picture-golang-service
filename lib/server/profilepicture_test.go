package server_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"example.com/lib/controllers"
	"example.com/lib/server"
	"github.com/gavv/httpexpect"
	"github.com/spf13/viper"
)

var testServer *httptest.Server

func setupTest() {
	viper.Set("bucket", "../../test-bucket")
	viper.Set("storagesystem", "mocked")
	router := server.NewRouter()
	testServer = httptest.NewServer(router)
}

func tearDownTest() {
	testServer.Close()
}

func TestProfilePictureServer(t *testing.T) {
	setupTest()
	testGuid := "test-guid"
	e := httpexpect.New(t, testServer.URL)

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
	tearDownTest()
}

func TestProfilePictureGet_PostiveCase(t *testing.T) {
	setupTest()
	e := httpexpect.New(t, testServer.URL)
	e.GET(server.PROFILE_PICTURE_PATH).WithHeader(controllers.GUID_HEADER, "sampleguid").
		Expect().
		Status(http.StatusOK)

	tearDownTest()
}

func TestProfilePictureGet_NegativeCase(t *testing.T) {
	setupTest()
	e := httpexpect.New(t, testServer.URL)
	e.GET(server.PROFILE_PICTURE_PATH).WithHeader(controllers.GUID_HEADER, "sampleguid-1").
		Expect().
		Status(http.StatusInternalServerError)

	tearDownTest()
}

func TestProfilePicturePost_PostiveCase(t *testing.T) {
	setupTest()
	e := httpexpect.New(t, testServer.URL)
	e.POST(server.PROFILE_PICTURE_PATH).WithHeader(controllers.GUID_HEADER, "sampleguid-post").
		WithMultipart().
		WithFile("photo", "../../test-bucket/example.jpg").
		Expect().
		Status(http.StatusOK)
	if _, err := os.Stat("../../test-bucket/sampleguid-post.jpe"); err == os.ErrNotExist {
		// file wasnot added, fail the test case
		t.Error("file was not saved")
		t.Fail()
	}

	tearDownTest()
}

func TestProfilePicturePost_NegativeCase(t *testing.T) {
	setupTest()
	e := httpexpect.New(t, testServer.URL)
	e.POST(server.PROFILE_PICTURE_PATH).WithHeader(controllers.GUID_HEADER, "sampleguid-post").
		WithMultipart().
		WithFile("somekey", "../../test-bucket/example.jpg").
		Expect().
		Status(http.StatusInternalServerError)
	tearDownTest()
}

func TestProfilePictureDelete_NegativeCase(t *testing.T) {
	setupTest()
	e := httpexpect.New(t, testServer.URL)

	e.DELETE(server.PROFILE_PICTURE_PATH).WithHeader(controllers.GUID_HEADER, "some-random-guid").
		Expect().
		Status(http.StatusInternalServerError)
	tearDownTest()
}
