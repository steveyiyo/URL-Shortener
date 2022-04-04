package cache

import (
	"fmt"
	"testing"
	"time"
)

// It's a test function.
func TestMain(t *testing.T) {
	InitRedis("127.0.0.1", "")
	AddData("hi", "pong", 5)
	status, data := QueryData("hi")
	if status {
		fmt.Println(data)
	} else {
		fmt.Println("QaQ")
	}
	time.Sleep(6 * time.Second)
	_, data = QueryData("hi")
	fmt.Println(data)
}
