package common

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

/*
 * @desc    字符串工具类
 * @author 	Zack
 * @date 	2020/9/10 18:08
 ****************************************
 */

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// 高效率获取随机字符串
func RandomString(n int) string {
	b := make([]byte, n)
	src := rand.NewSource(time.Now().UnixNano())
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// 返回 s到e范围内的数字，n位数的随机验证码
func RandomCode(s, e, n int) string {
	buffer := bytes.Buffer{}
	rand.Seed(time.Now().UnixNano())
	for i := s; i < n; i++ {
		buffer.WriteString(fmt.Sprintf("%d", rand.Intn(e)))
	}
	return buffer.String()
}

// 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// 获取数据，如果为空则调用函数
func StructValueDef(value interface{}, defValue string, callFunc func() string) string {
	vi := reflect.ValueOf(value)
	if vi.Kind() == reflect.Ptr {
		if vi.IsNil() {
			return defValue
		}
	}
	return callFunc()
}

// 判断是否为空，否则返回默认值
func GetStringDef(value string, defValue string) string {
	if value != "" {
		return value
	}
	return defValue
}
