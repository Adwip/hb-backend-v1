package library

import "time"

type TimeLib struct {
	currentTime    time.Time
	dateTimeFormat string
	dateFormat     string
	timeFormat     string
	dbTimeFormat   string
}

func Time() *TimeLib {
	time := &TimeLib{
		currentTime:    time.Now(),
		dateTimeFormat: "15:04:05 02-01-2006",
		dateFormat:     "02-01-2006",
		timeFormat:     "15:04:05",
		dbTimeFormat:   "02-01-2006 15:04:05",
	}
	return time
}

func (TimeLib) StringToTime() {

}

func (TimeLib) TimeToString() {

}

func (ct TimeLib) StringTimeNow() string {
	result := ct.currentTime.Format(ct.timeFormat)
	return result
}

func (ct TimeLib) DbTimeNow() string {
	result := ct.currentTime.Format(ct.dbTimeFormat)
	return result
}
