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

func AddURL(c *gin.Context) {
	var data Data
	c.BindJSON(&data)
	var return_data URLid
	timestampcheck, timestamp := Tools.ConvertTimetoUnix(data.EXPIREAT)
	if Tools.CheckLinkValid(data.URL) && (timestampcheck) {
		ShortID := Tools.RandomString(5)
		Database.AddData(ShortID, data.URL, timestamp)
		return_data = URLid{ID: ShortID, ShortURL: URL + ShortID}
		c.JSON(200, return_data)
	} else {
		return_data = URLid{ID: "", ShortURL: ""}
		c.JSON(400, return_data)
	}
}

func RedirectURL(c *gin.Context) {
	ID := c.Param("ShortID")
	// Search Link in DB
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
	route := gin.New()
	route.Use(gin.Logger(), gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	route.Use(cors.New(config))
	route.GET("/:ShortID", RedirectURL)
	route.POST("/api/v1/urls", AddURL)

	route.Run(Listen + ":" + Port)
}
