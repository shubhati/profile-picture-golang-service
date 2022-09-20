package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/lib/server"
	"github.com/spf13/viper"
)

/*
	init function is called before main function automatically by go compiler.
	init function here is reading the config file from ./resource/config.json file.
	if the config is not found or there is any error while reading the config, the program will
	halt.
*/
func init() {
	viper.AddConfigPath("./resources/")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found")
		} else {
			log.Fatal("Error in reading config file:", err.Error())
		}
	}

}

/*
	main function starts an http server which serves GET method on "/v1/profile-pic" endpoint.
*/
func main() {
	router := server.NewRouter()
	http.ListenAndServe(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")), router)
	// router.Run(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")))
}
