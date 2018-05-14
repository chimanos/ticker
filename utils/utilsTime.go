package utils

import (
	"time"
	"fmt"
)

func GetDateFromTimestamp(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

func GetTimestampFromDate(date string) int64 {
	layout := "01/02/2006 3:04:05 PM"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	return t.Unix()
}