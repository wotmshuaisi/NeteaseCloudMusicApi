// @APIVersion 1.0.0
// @Title NeteaseCloudMusic Api
// @Description 用golang写的网易云音乐API,请多多指教
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/ActingCute/NeteaseCloudMusicApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
