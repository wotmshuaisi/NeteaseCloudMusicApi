package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type MusicController struct {
	baseController
}


// @Title lyric
// @Description get music lyric
// @Success 200 {string} get music lyric success
// @router /lyric [get]
// @Param id query string true "The music id for lyric"
func (this *MusicController)Lyric() {
	this.NeedAPIinput("id")
	id, _ := this.GetInt64("id")
	if id < 1 {
		this.SetReturnData(500, "id error", nil)
		return
	}
	idstr := this.GetString("id")
	cname := idstr + "musiclyric"
	list := GetCache(cname)
	beego.Debug(list)
	if len(list) > 0 {
		this.SetReturnData(200, "ok", list)
		return
	}
	var detail interface{}
	detaildata, err := json.Marshal(detail)
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	body, err := this.Http(BaseApi  + MusicLyricApi + idstr, detaildata, "POST")
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	SetCache(cname, string(body), 600)
	this.SetReturnData(200, "ok", body)
}