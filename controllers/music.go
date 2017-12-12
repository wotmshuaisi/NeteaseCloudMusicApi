package controllers

import (
	"encoding/json"
)

type MusicController struct {
	BaseController
}

type Lyric struct {
	Sgc       bool `json:"sgc"`
	Sfy       bool `json:"sfy"`
	Qfy       bool `json:"qfy"`
	LyricUser LyricUserData `json:"lyricUser"`
	Lrc       LrcDta `json:"lrc"`
	Code      int `json:"code"`
}

type LyricUserData struct {
	Id       int64 `json:"id"`
	Status   int `json"status"`
	Demand   int `json:"demand"`
	Userid   int64 `json:"userid"`
	Nickname string `json:"nickname"`
	Uptime   int64 `json:"uptime"`
}

type LrcDta struct {
	Version int  `json:"version"`
	Lyric   string `json:"lyric"`
}

type Music struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Url    string `json:"url"`
	Pic    string `json:"pic"`
	Lrc    string `json:"lrc"`
	Id     int64  `json:"id"`
}

// @Title get music lyric
// @Description 获取歌词
// @Success 200 {object} controllers.Lyric
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
	//beego.Debug(list)
	if len(list) > 0 {
		var lrc Lyric
		json.Unmarshal([]byte(list), &lrc)
		this.SetReturnData(lrc.Code, "ok", lrc)
		return
	}
	var detail interface{}
	detaildata, err := json.Marshal(detail)
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	body, err := this.Http(BaseApi + MusicLyricApi + idstr, detaildata, "GET")
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	if len(string(body)) > 0 {
		SetCache(cname, string(body), 600)
	}
	var lrc Lyric
	json.Unmarshal(body, &lrc)
	this.SetReturnData(lrc.Code, "ok", lrc)
}