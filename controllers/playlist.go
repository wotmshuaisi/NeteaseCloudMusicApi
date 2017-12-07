package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

// Operations about PlayList
type PlayListController struct {
	baseController
}

type PlaylistDetail struct {
	Id        int64 `json:"id"`
	Offset    int `json:"offset"`
	Total     bool `json:"total"`
	Limit     int `json:"limit"`
	N         int        `json:"n"`
	CsrfToken string `json:"csrf_token"`
}

type Playlist struct {

}

// @Title detail
// @Description get user playlist detail
// @Success 200 {string} get user playlist detail success
// @router /detail [get]
// @Param id query string true "The id for playlist detail"
func (this *PlayListController)Detail() {
	this.NeedAPIinput("id")
	id, _ := this.GetInt64("id")
	if id < 1 {
		this.SetReturnData(500, "id error", nil)
		return
	}
	idstr := this.GetString("id")
	cname := idstr + "playlistdetail"
	list := GetCache(cname)
	beego.Debug(list)
	if len(list) > 0 {
		this.SetReturnData(200, "ok", list)
		return
	}
	var detail PlaylistDetail
	detail.CsrfToken = ""
	detail.Id = id
	detail.Limit = 1000
	detail.N = 1000
	detail.Total = true
	detail.Offset = 0
	detaildata, err := json.Marshal(detail)
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	body, err := this.Http(BaseApi + V3Dir + PlaylistDetailApi, detaildata, "POST")
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	SetCache(cname, string(body), 600)
	this.SetReturnData(200, "ok", body)
}