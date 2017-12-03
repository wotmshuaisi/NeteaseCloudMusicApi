package controllers

import (
	"github.com/ActingCute/NeteaseCloudMusicApi/models"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego"
	"strconv"
)

var (
	Cache redis.Conn
	RedisIp, RedisPort, RedisNetwork string
	RedisTime int64
)

func InitRedis() redis.Conn {
	if RedisEnable {
		url := RedisIp + ":" + RedisPort
		beego.Debug(url)
		c, err := redis.Dial(RedisNetwork, url)
		if err != nil {
			beego.Error(err)
		}
		return c
	}
	return nil
}
func init() {
	RedisEnable = models.GetAppConfBool("redis::enable")
	RedisIp = models.GetAppConf("redis::ip")
	RedisPort = models.GetAppConf("redis::port")
	RedisNetwork = models.GetAppConf("redis::network")
	time := models.GetAppConf("redis::time")
	RedisTime ,_ = strconv.ParseInt(time,10,64)
	if RedisTime < 600 {
		RedisTime = 600
	}
	if len(RedisIp) < 1 || len(RedisPort) < 1 || len(RedisNetwork) < 1 {
		RedisEnable = false
		return
	}
}

//歌单 uid+playlist
func SetCache(name string, value string, time int64) {
	Cache = InitRedis()
	if Cache == nil {
		beego.Info("SetCache Redis Close")
		return
	}
	_, err := Cache.Do("SET", name, value, "EX", time)
	if err != nil {
		beego.Error("redis set failed:", err)
	}
	defer Cache.Close()
}

func GetCache(name string) string {
	Cache = InitRedis()
	if Cache == nil {
		beego.Info("SetCache Redis Close")
		return ""
	}
	value, err := redis.String(Cache.Do("GET", name))
	if err != nil {
		beego.Error("redis get failed:", err)
	} else {
		beego.Error("Got GetCache %v \n", value)
	}
	defer Cache.Close()
	return value
}