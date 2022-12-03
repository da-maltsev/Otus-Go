package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	const (
		formatTime = "2006-01-02 15:04:05"
		timeHost   = "0.beevik-ntp.pool.ntp.org"
	)

	sharpTime, err := ntp.Time(timeHost)
	if err != nil {
		log.Fatal(err)
	}

	nowTime := time.Now()

	fmt.Printf("sharp time is: %s\n", sharpTime.Format(formatTime))
	fmt.Printf("now time is: %s\n", nowTime.Format(formatTime))
}
