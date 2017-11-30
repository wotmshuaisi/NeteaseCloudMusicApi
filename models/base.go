package models

import (
	"github.com/astaxie/beego"
)

func GetAppConf(name string) string {
	return beego.AppConfig.String(name)
}
