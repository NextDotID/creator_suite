package util

import "time"

// Datetime2DateString sql.datetime -> yyyy-mm-dd HH:MM:SS
func Datetime2DateString(datetime time.Time) string {
	timestamp := datetime.Unix()
	timeLayout := "2006-01-02 15:04:05"
	datetimeString := time.Unix(timestamp, 0).Format(timeLayout)
	return datetimeString
}

// DateString2oDatetime yyyy-mm-dd HH:MM:SS -> sql.datetime
func DateString2oDatetime(datetime string) (time.Time, error) {
	timeLayout := "2006-01-02 15:04:05"[0:MinInt(len(datetime), 19)]
	loc, _ := time.LoadLocation("Local")
	tmp, err := time.ParseInLocation(timeLayout, "2021-04-25 15:49:00", loc)
	return tmp, err
}
