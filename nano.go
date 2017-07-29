package main

import (
	"fmt"
	"time"
)

func nano() string {
	now := time.Now()
	now_rfc3339 := now.Format(time.RFC3339Nano)

	return now_rfc3339
}

func main() {
	n := nano()

	fmt.Println(n)
}
