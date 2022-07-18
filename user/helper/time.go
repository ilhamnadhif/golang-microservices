package helper

import (
	"time"
)

const TimeFormatSQL = "2006-01-02 15:04:05"

func EpochTimeToTime(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

func DateTimeParse(value string) time.Time {
	layoutFormat := "2006-01-02"
	date, err := time.Parse(layoutFormat, value)
	if err != nil {
		panic(err.Error())
	}
	return date
}
