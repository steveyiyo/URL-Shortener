package main

import (
	"log"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/steveyiyo/url-shortener/internal/Database"
	"github.com/steveyiyo/url-shortener/internal/Tools"
)

// Predefined variable and struct
type Config struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

var Listen string
var Host string
var Port string
var URL string

type Data struct {
	URL      string `json:"url"`
	EXPIREAT string `json:"expireAt"`
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
	timestampcheck, timestamp := Tools.ConvertTimetoUnix(data.EXPIREAT)

	// Check Link and Time Valid
	if Tools.CheckLinkValid(data.URL) && (timestampcheck) {

		// Random Short ID
		ShortID := Tools.RandomString(5)

		// Add data to DB
		Database.AddData(ShortID, data.URL, timestamp)

		// Return result
		return_data = URLid{ID: ShortID, ShortURL: URL + ShortID}
		c.JSON(200, return_data)
	} else {

		// Return result
		return_data = URLid{ID: "", ShortURL: ""}
		c.JSON(400, return_data)
	}
}

// RedirectURL
func RedirectURL(c *gin.Context) {

	// Get Short ID from URL
	ID := c.Param("ShortID")

	// Query Link in DB
	Check, Link := Database.QueryData(ID)
	if Check {
		c.Redirect(301, Link)
	} else {
		c.Status(404)
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

	// Init Database
	Database.CreateTable()

	// Init Web Server
	route := gin.New()
	route.Use(gin.Logger(), gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	route.Use(cors.New(config))
	route.GET("/:ShortID", RedirectURL)
	route.POST("/api/v1/urls", AddURL)

	route.Run(Listen + ":" + Port)
}
