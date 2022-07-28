package utils

import "time"
import "fmt"

type TimeLib struct {
	userDateFormat string
	dateTimeFormat string
	dateFormat     string
	timeFormat     string
	dbTimeFormat   string
}

func Time() *TimeLib {
	time := &TimeLib{
		userDateFormat: "02-01-2006 15:04:05",
		dateTimeFormat: "15:04:05 02-01-2006",
		dateFormat:     "02-01-2006",
		timeFormat:     "15:04:05",
		dbTimeFormat:   "2006-01-02 15:04:05",
	}
	return time
}

func (tl TimeLib) StringToTime() {

}

func (tl TimeLib) TimeToString() string {
	result := ""
	return result
}

func (ct TimeLib) CurrentTimeString() string {
	result := time.Now().Format(ct.timeFormat)
	return result
}

func (TimeLib) CurrentTimeUnix() int64 {
	result := time.Now().Unix()
	return result
}

func (TimeLib) CurrentTimeUTC() time.Time {
	result := time.Now().UTC()
	return result
}

func (ct TimeLib) CurrentDateTimeDbFormat() string {
	result := time.Now().UTC().Format(ct.dbTimeFormat)
	return result
}

func (t TimeLib) ConvertToUTC(stringTime string, timezone string) (string, error) {
	loc, errLoc := time.LoadLocation(timezone)
	if errLoc != nil {
		return "", errLoc
	}
	result, errParse := time.ParseInLocation(t.userDateFormat, stringTime, loc)
	if errParse != nil {
		return "", errParse
	}
	fmt.Println(result)
	return result.UTC().Format(t.dbTimeFormat), nil
}
