package common

import "time"

/*
 * @desc    日期时间工具类
 * @author 	Zack
 * @date 	2020/9/29 15:55
 ****************************************
 */
var (
	// go中的时间格式化必须使用这个时间 2006-01-02 15:04:05
	YYYYMMdd_HHmmss = DateFormatLayout{"2006-01-02 15:04:05"}
	YYYYMMddHHmmss  = DateFormatLayout{"20060102150405"}
	YYYYMMdd_HHmm   = DateFormatLayout{"2006-01-02 15:04"}
	YYYYMMdd        = DateFormatLayout{"2006-01-02"}
	YYYYMM          = DateFormatLayout{"2006-01"}
)

type DateFormatLayout struct {
	Name string
}

func DateToStringDefault(time time.Time) string {
	return DateToString(time, YYYYMMdd_HHmmss)
}
func DateToString(time time.Time, layout DateFormatLayout) string {
	return time.Format(layout.Name)
}
func StringToDateDefault(timeStr string) time.Time {
	return StringToDate(timeStr, YYYYMMdd_HHmmss)
}
func StringToDate(timeStr string, layout DateFormatLayout) time.Time {
	parse, _ := time.Parse(layout.Name, timeStr)
	return parse
}
