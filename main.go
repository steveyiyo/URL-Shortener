package main

import (
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/steveyiyo/url-shortener/internal/cache"
	Database "github.com/steveyiyo/url-shortener/internal/database"
	Tools "github.com/steveyiyo/url-shortener/internal/tools"
)

// Predefined variable and struct
type Config struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type Result struct {
	Status  bool
	Message string
}

var Listen string
var Host string
var Port string
var URL string

type Data struct {
	URL       string `json:"url"`
	ExpiredAt string `json:"expireAt"`
}

type URLid struct {
	ID       string `json:"id"`
	ShortURL string `json:"shortUrl"`
}

// AddURL
func AddURL(c *gin.Context) {
	// Get JSON Data
	var data Data
	c.BindJSON(&data)

	// Init return result
	var return_data URLid

	// Check Time and Convert to Unix format
	isTimestampOk, timestamp := Tools.ConvertTimetoUnix(data.ExpiredAt)

	// Check Limit IP
	limit_check, _ := cache.QueryData(c.ClientIP())

	// Check Limit
	if !limit_check {
		// Check Link and Time Valid
		if Tools.CheckLinkValid(data.URL) && (isTimestampOk) {

			// Random Short ID
			ShortID := Tools.RandomString(5)

			// Add data to DB
			Database.AddData(ShortID, data.URL, timestamp)

			// Add Limit IP to Redis
			cache.AddData(c.ClientIP(), "", 5)

			// Return result
			return_data = URLid{ID: ShortID, ShortURL: URL + ShortID}
			c.JSON(200, return_data)
		} else {
			// Return result
			return_data = URLid{ID: "", ShortURL: ""}
			c.JSON(400, return_data)
		}
	} else {
		// Return result
		var return_result Result
		return_result = Result{Status: false, Message: "Too many requests, please try again later."}
		c.JSON(400, return_result)
	}
}

// RedirectURL
func RedirectURL(c *gin.Context) {

	// Get Short ID from URL
	ID := c.Param("ShortID")

	// Query Link in Redis
	isExist, URL := cache.QueryData(ID)
	if !isExist {
		// Query Link in DB
		Check, Link := Database.QueryData(ID)
		if Check {
			// Add hit to Redis
			cache.AddData(ID, Link, 30)
			c.Redirect(301, Link)
		} else {
			// Add miss link to Redis (Not Found or Expire)
			cache.AddData(ID, "MISS", 30)

			// Return result
			var return_result Result
			return_result = Result{Status: false, Message: "Not Found."}
			c.JSON(404, return_result)
		}
	} else {
		if URL == "MISS" {
			// Return result
			var return_result Result
			return_result = Result{Status: false, Message: "Not Found."}
			c.JSON(404, return_result)
		} else {
			c.Redirect(301, URL)
		}
	}
}

func main() {
	// Load Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); nil != err {
		log.Fatalf("Failed to read config: %v\n", err)
	}

	// Define Config
	Listen = viper.GetString("Listen")
	Host = viper.GetString("Host")
	Port = strconv.Itoa(viper.GetInt("Port"))
	URL = Host + ":" + Port + "/"

	// Check Config
	// Tools.CheckIPAddress(Listen)

	// Init Redis
	cache.InitRedis()

	// Init Database
	Database.CreateTable()

	// Init Web Server
	gin.SetMode(gin.ReleaseMode)
	route := gin.New()
	route.Use(gin.Logger(), gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	route.Use(cors.New(config))
	route.GET("/:ShortID", RedirectURL)
	route.POST("/api/v1/urls", AddURL)

	route.Run(Listen + ":" + Port)
}
