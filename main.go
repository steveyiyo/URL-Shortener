package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/dcard-url-shortener/internal/Database"
	"github.com/steveyiyo/dcard-url-shortener/internal/Tools"
)

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
		return_data = URLid{ID: ShortID, ShortURL: "http://localhost:19247/" + ShortID}
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
	Database.CreateTable()
	route := gin.New()
	route.Use(gin.Logger(), gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	route.Use(cors.New(config))
	route.GET("/:ShortID", RedirectURL)
	route.POST("/api/v1/urls", AddURL)

	route.Run("127.0.0.1:19247")
}
