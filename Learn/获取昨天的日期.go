package main

import (
	"log"
	"time"
)

func main() {
	timenow := time.Now().AddDate(0, 0, -1).Format("20060102")
	log.Printf(timenow)
}
