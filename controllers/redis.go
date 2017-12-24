package controllers

import (
	"github.com/ActingCute/blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

var (
	Cache                            cache.Cache
	RedisIp, RedisPort, RedisNetwork string
	RedisTime                        int64
	RedisEnable                      bool
)

func init() {
	//redis
	RedisEnable = models.GetAppConfBool("redis::enable")
	RedisIp = models.GetAppConf("redis::ip")
	RedisPort = ":"+ models.GetAppConf("redis::port")
	RedisNetwork = models.GetAppConf("redis::network")
	time := models.GetAppConf("redis::time")
	RedisTime, _ = strconv.ParseInt(time, 10, 64)
	if RedisTime < 600 {
		RedisTime = 600
	}
	if len(RedisIp) < 1 || len(RedisPort) < 1 || len(RedisNetwork) < 1 {
		RedisEnable = false
		return
	}
	beego.Debug("{\"conn\": \""+RedisIp+RedisPort+"\"}")
	var err error
	Cache, err = cache.NewCache("redis", "{\"conn\": \""+RedisIp+RedisPort+"\"}")
	if err != nil {
		RedisEnable = false
		beego.Error("init Cache err,err:", err)
	}else{
		beego.Info("init Cache Success")
	}
}

func GetCacheInt(name string) int {
	if !RedisEnable {
		beego.Info("GetCacheInt RedisEnable is ", RedisEnable)
		return 0
	}
	var err error
	val, _ := redis.Int(Cache.Get(name), err)
	return val
	return 0
}

func GetCache(name string) string {
	if !RedisEnable {
		beego.Info("GetCacheString RedisEnable is ", RedisEnable)
		return ""
	}
	var err error
	val, _ := redis.String(Cache.Get(name), err)
	return val
	return ""
}

func SetCacheInt(name string, val int, timeout int64) {
	if !RedisEnable {
		beego.Info("SetCacheInt RedisEnable is ", RedisEnable)
		return
	}
	err := Cache.Put(name, val, time.Duration(timeout)*time.Second)

	if nil != err {
		beego.Error("SetCacheInt error. name:", name, " err:", err)
		return
	}

	beego.Debug("SetCacheInt is ok. name:", name)
}

func SetCache(name string, val string, timeout int64) {
	beego.Debug("timeout=",time.Duration(timeout))
	if !RedisEnable {
		beego.Info("SetCacheString RedisEnable is ", RedisEnable)
		return
	}
	err := Cache.Put(name, val, time.Duration(timeout)*time.Second)

	if nil != err {
		beego.Error("SetCacheString error. name:", name, " err:", err)
		return
	}

	beego.Debug("SetCacheString is ok. name:", name)
}

func SetCacheIncr(name string) {
	if !RedisEnable {
		beego.Info("SetCacheIncr RedisEnable is ", RedisEnable)
		return
	}
	err := Cache.Incr(name)
	if nil != err {
		beego.Error("Incr Cache error. name:", name, " err:", err)
		return
	}
	beego.Debug("Incr Cache is ok. name:", name)
}
