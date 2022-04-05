# URL Shortener

A URL Shortener.

## 3rd Part Library Usage

- Web Engineering
  - Use gin for web framework.
  - Use cors for cross-domain requests. It will let you make requests from any domain.
- Config file
  - Use viper to parse the config file, to make the codes more flexible.
- Database
  - Use sqlite3 for database.
- Cache Storage
  - Use Redis as cache storage.
  - If the Short ID does not exist in the DB, it will add a miss to the Redis.
  - If the Short ID exists in the DB, it will add as hit to the Redis.
  - Limit IP Address.

For the full list, please check on [go.mod](go.mod).

## Features

- Init
  - Check if the table or file exists. If not, create it.
  - Loading .env file.
- Add Short Link
  - HTTP POST with JSON body.
  - Check if the input (url, expireAt) is valid.
    - `url`: require a valid URL. Use `url.ParseRequestURI` to check.
    - `expireAt`: require a valid date. Use `tools.ConvertTimetoUnix` to check. (btw, It will be converted to unix timestamp.)
- Redirect
  - HTTP GET with URL query.
  - It will query the database to get the origin URL. (btw, It will also add a hit to the Redis.)

### Function List

```go
func AddURL(c *gin.Context)               // Request for a Short Link
func RedirectURL(c *gin.Context)          // Redirect to the origin link

func CreateTable()                        // Init Table (If a table or file is not found, it will create one.)
func AddData(ShortID string, Link string, ExpireAt int64) // Prepend to DB
func QueryData(ID string) (bool, string)  // Query the Origin Link
func ErrCheck(err error)                  // Check Error

func RandomString(length int) string      // Generate a 5 char string
func CheckLinkValid(Link string) bool     // Check if Link valid
func ConvertTimetoUnix(date string) (bool, int64) // Convert the time format from RFC 3339 UTC to Unix
func CheckIPAddress(ip string) bool       // Check if IP Address valid
```

## Install

- Download the libraries  
`go mod download`

- Test the server  
`go run main.go`

- Build  
`go build -o url_shortener`

## Usage

Download the [release version](https://github.com/steveyiyo/URL-Shortener/releases) from here.

Upload URL API:

1. "url" require a valid URL. (Eg: https://google.com/)
2. "expireAt" require RFC 3399 UTC format.

```
curl -X POST -H "Content-Type:application/json" http://localhost/api/v1/urls -d '{
    "url": "https://steveyi.net/",
    "expireAt": "2021-02-08T09:20:41Z"
}'
```

Response:

1. "id" will generate a 5 character random string.
2. "shortUrl" will generate a short URL

```
{
    "id": "<url_id>",
    "shortUrl": "http://localhost/<url_id>"
}
```

Redirect URL API:

```
curl -L -X GET http://localhost/<url_id> => REDIRECT to original URL # HTTP 301
```

## Run with the docker?

- Download the docker image
`docker pull steveyiyo/url-shortener:latest`

- Run the docker image
`sudo docker run -p 0.0.0.0:7600:7600 steveyiyo/url-shortener`

## TODO

- [] docker-compose.yaml