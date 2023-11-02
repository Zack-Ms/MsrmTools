package common

import (
	"github.com/patrickmn/go-cache"
	"time"
)

/*
 * @desc    cache 工具类
 * @author 	zk
 * @date 	2021/3/29 18:28
 ****************************************
 */
var (
	cacheUtil *cache.Cache
)

func GetCache() *cache.Cache {
	if cacheUtil == nil {
		// 创建一个cache对象，默认ttl 5分钟，每10分钟对过期数据进行一次清理
		cacheUtil = cache.New(5*time.Minute, 10*time.Minute)
	}
	return cacheUtil
}

//
func Incr(key string, value, n int, exp time.Duration) (int, error) {
	cacheClient := GetCache()
	incr, get := cacheClient.Get(key)
	if get {
		incr, _ = cacheClient.IncrementInt(key, n)
	} else {
		cacheClient.Set(key, value, exp)
		incr = 1
	}
	return incr.(int), nil
}
