package models

import (
	"github.com/astaxie/beego"
)

func GetAppConf(name string) string {
	return beego.AppConfig.String(name)
}

func GetAppConfBool(name string) bool {
	b, err := beego.AppConfig.Bool(name)
	if err != nil {
		return false
	}
	return b
}