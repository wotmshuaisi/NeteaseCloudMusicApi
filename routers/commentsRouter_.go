package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/ActingCute/NeteaseCloudMusicApi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ActingCute/NeteaseCloudMusicApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "CellphoneLogin",
			Router: `/cellphonelogin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ActingCute/NeteaseCloudMusicApi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ActingCute/NeteaseCloudMusicApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Playlist",
			Router: `/playlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/ActingCute/NeteaseCloudMusicApi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ActingCute/NeteaseCloudMusicApi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/detail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
