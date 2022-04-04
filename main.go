package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/steveyiyo/url-shortener/internal/cache"
	"github.com/steveyiyo/url-shortener/internal/database"
	"github.com/steveyiyo/url-shortener/internal/webserver"
	"github.com/steveyiyo/url-shortener/package/tools"
)

func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); nil != err {
		log.Fatalf("Failed to read config: %v\n", err)
	}

	// Define Config
	Listen := viper.GetString("Listen")
	Host := viper.GetString("Host")
	Port := strconv.Itoa(viper.GetInt("Port"))
	URL := Host + ":" + Port + "/"
	Redis_Addr := os.Getenv("Redis_Addr")
	Redis_Pwd := os.Getenv("Redis_Pwd")

	// Check Config
	if tools.CheckIPAddress(Listen) {
		// Init Redis
		cache.InitRedis(Redis_Addr, Redis_Pwd)

		// Init Database
		database.Init()

		// Init Web Server
		webserver.Init(Listen, Host, Port, URL)
	} else {
		log.Fatal("Error: Listen IP Address is not valid")
	}
}
