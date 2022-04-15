package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/steveyiyo/url-shortener/internal/cache"
	"github.com/steveyiyo/url-shortener/internal/database"
	"github.com/steveyiyo/url-shortener/internal/webserver"
	"github.com/steveyiyo/url-shortener/pkg/tools"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Define Config
	Listen := os.Getenv("Listen")
	Host := os.Getenv("Host")
	Port := os.Getenv("Port")
	URL := Host + ":" + Port + "/"
	Redis_Addr := os.Getenv("Redis_Addr")
	Redis_Pwd := os.Getenv("Redis_Pwd")

	fmt.Println(Redis_Addr)

	if Listen == "" || Host == "" || Port == "" || Redis_Addr == "" {
		log.Fatal("Error to loading environment.")
	}

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
