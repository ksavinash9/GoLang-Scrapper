package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {

	router := NewRouter()

	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
  if err != nil {
		log.Fatal("Config file not found")
  }

	log.Fatal(http.ListenAndServe(":8080", router))
}
