/**
  create by yy on 2019-07-26
*/

package go_helper

import (
	"time"
)

// 得到当前时间戳(unix 纪元)
func GetNowTimeStamp() int64 {
	return time.Now().Unix()
}

// 根据时间戳获取 日期 返回一个字符串
// According to the time stamp, get a date string.
// format is 2006-01-02 15:04:05
func GetDateByTimeStamp(timeStamp int64) string {
	if timeStamp == 0 {
		timeStamp = time.Now().Unix()
	}
	return time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
}

// 根据时间戳和想要的时间格式进行获取时间
func GetDateByTimeStampAndFormat(timeStamp int64, format string) string {
	if timeStamp == 0 {
		timeStamp = time.Now().Unix()
	}
	return time.Unix(timeStamp, 0).Format(format)
}

// 获取当前时间
func GetNowDate() string {
	return time.Unix(GetNowTimeStamp(), 0).Format("2006-01-02 15:04:05")
}

// 获取当前时间(根据给定的时间格式)
func GetNowDateByFormat(format string) string {
	return time.Unix(GetNowTimeStamp(), 0).Format(format)
}

// 根据时间戳获取 日期 返回一个字符串
// According to the time stamp, get a date string.
// format is 2006-01-02 15:04:05
func GetDateByTimeStampUTC(timeStamp int64) string {
	if timeStamp == 0 {
		timeStamp = time.Now().Unix()
	}
	return time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
}

// 根据时间戳和想要的时间格式进行获取时间
func GetDateByTimeStampAndFormatUTC(timeStamp int64, format string) string {
	if timeStamp == 0 {
		timeStamp = time.Now().UTC().Unix()
	}
	return time.Unix(timeStamp, 0).UTC().Format(format)
}

// 获取当前时间
func GetNowDateUTC() string {
	return time.Unix(GetNowTimeStamp(), 0).UTC().Format("2006-01-02 15:04:05")
}

// 获取当前时间(根据给定的格式)
func GetNowDateByFormatUTC(format string) string {
	return time.Unix(GetNowTimeStamp(), 0).UTC().Format(format)
}

// 根据日期获取时间戳
func GetStampByDateAndFormat(date string, format string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(format, date, loc)
	return theTime.Unix()
}

// 根据时间戳和给定格式以及 要加(减)的 年 月 日 获取对应时间字符串
func GetDateByStampAndReduceDate(stamp int64, format string, date ...int) string {
	return time.Unix(stamp, 0).AddDate(date[0], date[1], date[2]).Format(format)
}

func GetDateByStampAndReduceDateUTC(stamp int64, format string, date ...int) string {
	return time.Unix(stamp, 0).AddDate(date[0], date[1], date[2]).UTC().Format(format)
}
