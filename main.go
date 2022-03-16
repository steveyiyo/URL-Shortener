package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/steveyiyo/dcard-url-shortener/internal/Database"
	"github.com/steveyiyo/dcard-url-shortener/internal/Tools"
)

type Data struct {
	URL      string `json:"url"`
	EXPIREAT string `json:"expireAt"`
}

type URLid struct {
	ID       string `json:"id"`
	ShortURL string `json:"shortURL"`
}

func AddURL(c *gin.Context) {
	var data Data
	c.BindJSON(&data)
	var return_data URLid
	if Tools.CheckLinkValid(data.URL) && (data.EXPIREAT != "") {
		ShortID := Tools.RandomString(5)
		Database.AddData(ShortID, data.URL, data.EXPIREAT)
		return_data = URLid{ID: ShortID, ShortURL: "http://localhost:19247/api/v1/urls/" + ShortID}
		c.JSON(200, return_data)
		fmt.Println(time.Now().UTC().Format(data.EXPIREAT))
	} else {
		return_data = URLid{ID: "", ShortURL: ""}
		c.JSON(400, return_data)
	}
}

func QueryDB(ID string) string {
	// Search Link in DB
	Check, URL := Database.QueryData(ID)
	if Check {
		return URL
	} else {
		return ""
	}
}

func RedirectURL(c *gin.Context) {
	ID := c.Param("ShortID")
	// Search Link in DB
	Link := QueryDB(ID)
	if Link != "" {
		c.Redirect(301, Link)
	} else {
		c.Status(404)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := gin.New()
	route.Use(gin.Logger(), gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	route.Use(cors.New(config))
	route.GET("/:ShortID", RedirectURL)
	route.POST("/api/v1/urls", AddURL)

	route.Run("127.0.0.1:19247")
}
