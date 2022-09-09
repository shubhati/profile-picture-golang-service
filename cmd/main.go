package main

import (
	"example.com/lib/server"
)

const bucketName = "test-bucket"

/*
	main function starts an http server which serves GET method on "/v1/profile-pic" endpoint.
*/
func main() {
	router := server.NewRouter()
	router.Run("localhost:8087")
}
