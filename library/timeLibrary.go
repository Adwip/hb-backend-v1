package library

import "time"

type TimeLib struct {
	dateTimeFormat string
	dateFormat     string
	timeFormat     string
	dbTimeFormat   string
}

func Time() *TimeLib {
	time := &TimeLib{
		dateTimeFormat: "15:04:05 02-01-2006",
		dateFormat:     "02-01-2006",
		timeFormat:     "15:04:05",
		dbTimeFormat:   "2006-02-01 15:04:05",
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

func (TimeLib) CurrentTimeISO() time.Time {
	result := time.Now().UTC()
	return result
}

func (ct TimeLib) CurrentTImeDbFormat() string {
	result := time.Now().Format(ct.dbTimeFormat)
	return result
}
