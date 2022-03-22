# Technical details

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

For the full list, please check on [go.mod](go.mod).

## Features

- Init
  - Check if the table or file exists. If not, create it.
- Add Short Link
  - HTTP POST with JSON body.
  - Check if the input (url, expireAt) is valid.
    - `url`: require a valid URL. Use `url.ParseRequestURI` to check.
    - `expireAt`: require a valid date. Use `tools.ConvertTimetoUnix` to check. (btw, It will be converted to unix timestamp.)
- Redirect
  - HTTP GET with URL query.
  - It will query the database to get the origin URL.

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