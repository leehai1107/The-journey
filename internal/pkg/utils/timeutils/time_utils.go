package timeutils

import (
	"fmt"
	"sync"
	"time"
)

const (
	customRFC3339 = "2006-01-02T15:04:05"
	FormatYMD     = "yyyy-MM-dd"
	// YYYY_MM_DD_HH_MM_SS_SSS is the format "2006-01-02 15:04:05.000".
	YYYY_MM_DD_HH_MM_SS_SSS = "2006-01-02 15:04:05.000"
	// YYYY_MM_DD_HH_MM_SS is the format "2006-01-02 15:04:05".
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
	// YYYY_MM_DD is the format "2006-01-02".
	YYYY_MM_DD = "2006-01-02"
	// DD_MM_YYYY is the format "02-01-2006".
	DD_MM_YYYY = "02-01-2006"
	// DD_MM_YYYY_HH_MM_SS is the format "02-01-2006 15:04:05".
	DD_MM_YYYY_HH_MM_SS = "02-01-2006 15:04:05"
	// DD_MM_YYYY_HH_MM_SS_SSS is the format "02-01-2006 15:04:05.000".
	DD_MM_YYYY_HH_MM_SS_SSS = "02-01-2006 15:04:05.000"

	DMFormat = "d/m"
)

var (
	locationGMT07 *time.Location
	once          sync.Once
	initialized   bool
)

type NowFn func() time.Time
type NowTimestampFn func() int64

//nolint:gochecknoinits
func Init() {
	initTimezones()
}

func initTimezones() {
	once.Do(func() {
		var err error
		// Load required location
		locationGMT07, err = time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			panic(err)
		}
		initialized = true
	})
}

func GMT07Location() *time.Location {
	if !initialized {
		fmt.Println("Cannot use GMT+07 timezone, have you forgotten to call InitTimezones()?")
		return time.UTC
	}
	return locationGMT07
}

func NowInGMT07String(format string) string {
	return TimeInGMT07String(time.Now(), format)
}

func TimeInGMT07String(t time.Time, format string) string {
	result := t.In(GMT07Location()).Format(format)
	return result
}

func TimestampToGMT07Time(timestamp int64) time.Time {
	return ConvertTimeToGMT07(time.Unix(timestamp, 0))
}

func TimestampToTimeUTC(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).UTC()
}

func ConvertTimeToGMT07(t time.Time) time.Time {
	return t.In(GMT07Location())
}

func ConvertToUnixTime(t time.Time) int64 {
	return t.Unix()
}

/*
*
Convert timestamp to string with custom RFC3339 format without GMT
ex: 2006-01-02T15:04:05Z07:00
*/
func ConvertUnixTimeRFC3339String(timeStamp int64) string {
	return time.Unix(timeStamp, 0).In(GMT07Location()).Format(customRFC3339)
}

/*
*
Convert timestamp to string with custom RFC3339 format without GMT07
ex: 2006-01-02T15:04:05Z07:00
*/
func ParseStringToUnixTimestampLocation(timeStr string) int64 {
	t, err := time.ParseInLocation(customRFC3339, timeStr, GMT07Location())
	if err != nil {
		return 0
	}
	return t.Unix()
}

// ParseStringToTime convert string to time object
func ParseStringToTime(timeString string) time.Time {
	tm, _ := time.ParseInLocation(customRFC3339, timeString, GMT07Location())
	return tm
}

// TimeBeginDayByTime convert time to time begin of day
func TimeBeginDayByTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 00, 00, 00, 0, t.Location())
}

// TimeEndDayByTime convert time to time end of day
func TimeEndDayByTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

// GetDayOfWeekName get name of day of week, sunday is 0, 1 to 6 is from monday to saturday
func GetDayOfWeekNameNormalLetter(dayOfWeek time.Weekday) string {
	switch dayOfWeek {
	case time.Sunday:
		return "chủ nhật"
	case time.Monday:
		return "thứ 2"
	case time.Tuesday:
		return "thứ 3"
	case time.Wednesday:
		return "thứ 4"
	case time.Thursday:
		return "thứ 5"
	case time.Friday:
		return "thứ 6"
	case time.Saturday:
		return "thứ 7"
	}
	return ""
}

func GetDayOfWeekNameShort(dayOfWeek time.Weekday) string {
	switch dayOfWeek {
	case time.Sunday:
		return "CN"
	case time.Monday:
		return "T2"
	case time.Tuesday:
		return "T3"
	case time.Wednesday:
		return "T4"
	case time.Thursday:
		return "T5"
	case time.Friday:
		return "T6"
	case time.Saturday:
		return "T7"
	}
	return ""
}

// GetDayOfWeekName get name of day of week, sunday is 0, 1 to 6 is from monday to saturday
func GetDayOfWeekNameUpperFirstLetter(dayOfWeek time.Weekday) string {
	switch dayOfWeek {
	case time.Sunday:
		return "Chủ nhật"
	case time.Monday:
		return "Thứ 2"
	case time.Tuesday:
		return "Thứ 3"
	case time.Wednesday:
		return "Thứ 4"
	case time.Thursday:
		return "Thứ 5"
	case time.Friday:
		return "Thứ 6"
	case time.Saturday:
		return "Thứ 7"
	}
	return ""
}

func ParseStringDateToFormatDate(timeString string, format string) string {
	tm, _ := time.ParseInLocation(customRFC3339, timeString, GMT07Location())
	return getDateFormat(tm, format)
}

func ParseTimestampToFormatDate(timeStamp int64, format string) string {
	tm := time.Unix(timeStamp, 0).In(GMT07Location())
	return getDateFormat(tm, format)
}

func GetDaysBetweenDates(first, second time.Time) int64 {
	days := first.Sub(second).Hours() / 24
	return int64(days)
}

func getDateFormat(t time.Time, format string) string {
	switch format {
	case "d/m":
		return fmt.Sprintf("%d/%d", t.Day(), t.Month())
	case "d/m/yyyy":
		return fmt.Sprintf("%d/%d/%d", t.Day(), t.Month(), t.Year())
	case "dd/mm/yyyy":
		return fmt.Sprintf("%02d/%02d/%04d", t.Day(), t.Month(), t.Year())
	case "h:m d/m/yyyy":
		return fmt.Sprintf("%d:%d - %d/%d/%d", t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year())
	case "hh:mm d/m/yyyy":
		return fmt.Sprintf("%02d:%02d - %d/%d/%d", t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year())
	case "hh:mm dd/mm/yyyy":
		return fmt.Sprintf("%02d:%02d %02d/%02d/%d", t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year())
	case "mm/yyyy":
		return fmt.Sprintf("%02d/%d", t.Month(), t.Year())
	case "w (d/m)":
		return fmt.Sprintf("%s (%d/%d)", GetDayOfWeekNameNormalLetter(t.Weekday()), t.Day(), t.Month())
	case "hh:mm - d/m/yyyy":
		return fmt.Sprintf("%02d:%02d - %d/%d/%d", t.Hour(), t.Minute(), t.Day(), t.Month(), t.Year())
	case "hh:mm":
		return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
	}
	return fmt.Sprintf("%02d/%02d/%d", t.Day(), t.Month(), t.Year())
}

func ParseOpenTimeText(opensAt int64, closesAt int64) string {
	if opensAt == 0 || closesAt == 0 {
		return ""
	}

	timeOpensAt := time.Unix(opensAt, 0)
	timeClosesAt := time.Unix(closesAt, 0)

	// Build minute opens at
	textOpensAtMinute := fmt.Sprintf("%02d", timeOpensAt.Minute())
	// Build minute closes at
	textClosesAtMinute := fmt.Sprintf("%02d", timeClosesAt.Minute())
	return fmt.Sprintf("%v:%s - %v:%s", timeOpensAt.Hour(), textOpensAtMinute, timeClosesAt.Hour(), textClosesAtMinute)
}

func IsOnTheSameDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func NowInGMT07StringRFC3339() string {
	return TimeInGMT07String(time.Now(), customRFC3339)
}

func TimeInGMT07StringRFC3339(t time.Time) string {
	result := t.In(GMT07Location()).Format(customRFC3339)
	return result
}

func IsEqualDate(t1, t2 time.Time) bool {
	if t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day() {
		return true
	}

	return false
}

func GetBeginTimeOfDay(timestamp int64) int64 {
	t := time.Unix(timestamp, 0)
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Unix()
}
