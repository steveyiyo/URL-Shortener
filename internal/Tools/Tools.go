package Tools

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/url"
	"strings"
	"time"
)

// Generates a random string of a given length
func RandomString(length int) string {
	rand.Seed(time.Now().Unix())

	var output strings.Builder

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJULMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return (output.String())
}

// Check link valid
func CheckLinkValid(Link string) bool {
	check_status := false
	u, err := url.ParseRequestURI(Link)
	if err != nil {
		log.Println(err)
		log.Println(u)
		check_status = false
	} else {
		check_status = true
	}
	return check_status
}

// Check Time Valid and convert to Unix format
func ConvertTimetoUnix(date string) (bool, int64) {
	layout := "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, date)
	var status = true
	if err != nil {
		fmt.Println(err)
		status = false
	}
	return status, t.Unix()
}

// Check IP Valid
func CheckIPAddress(ip string) bool {
	var isCorrect bool
	if net.ParseIP(ip) == nil {
		isCorrect = false
	} else {
		isCorrect = true
	}
	return isCorrect
}

// Check Error
func ErrCheck(err error) bool {
	var check_status bool
	if err != nil {
		log.Println(err)
		check_status = false
	} else {
		check_status = true
	}
	return check_status
}
