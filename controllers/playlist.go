package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	//"time"
)

// Operations about PlayList
type PlayListController struct {
	BaseController
}

type PlaylistDetail struct {
	Id        int64 `json:"id"`
	Offset    int `json:"offset"`
	Total     bool `json:"total"`
	Limit     int `json:"limit"`
	N         int        `json:"n"`
	CsrfToken string `json:"csrf_token"`
}

type PlaylistJson struct {
	Playlist PlaylistData `json:"playlist"`
}

type PlaylistData struct {
	AdType                int `json:"adType"`
	CloudTrackCount       int `json:"cloudTrackCount"`
	CommentCount          int `json:"commentCount"`
	CommentThreadId       string `json:"commentThreadId"`
	CoverImgId            int64 `json:"coverImgId"`
	CoverImgUrl           string `json:"coverImgUrl"`
	CreateTime            int64 `json:"createTime"`
	Creator               CreatorData `json:"creator"`
	Description           []string `json:"description"`
	HighQuality           bool `json:"highQuality"`
	Id                    int64 `json:"id"`
	Name                  string `json:"name"`
	NewImported           bool `json:"newImported"`
	Ordered               bool `json:"ordered"`
	PlayCount             int `json:"playCount"`
	Privacy               int `json:"privacy"`
	ShareCount            int `json:"shareCount"`
	SpecialType           int `json:"specialType"`
	Status                int `json:"status"`
	Subscribed            bool `json:"subscribed"`
	SubscribedCount       int `json:"subscribedCount"`
	Subscribers           []string `json:"subscribers"`
	Tags                  []string `json:"tags"`
	TrackCount            int `json:"trackCount"`
	TrackIds              []TrackIdsData `json:"TrackIds"`
	TrackNumberUpdateTime int64 `json:"trackNumberUpdateTime"`
	TrackUpdateTime       int64 `json:"trackUpdateTime"`
	Tracks                []TracksData `json:"tracks"`
	UpdateTime            int64 `json:"updateTime"`
	UserId                int64 `json:"userId"`
}

type  TracksData struct {
	A           string `json:"a"`
	Al          AlData `json:"al"`
	Alia        []string `json:"alia"`
	Ar          []ArData `json:"ar"`
	Cd          int `json:"cd"`
	Cf          string `json:"cf"`
	Copyright   int `json:"copyright"`
	Cp          int `json:"cp"`
	Crbt        []string `json:"crbt"`
	DjId        int `json:"djId"`
	Dt          int `json:"dt"`
	Fee         int `json:"fee"`
	Ftype       int `json:"ftype"`
	H           HData `json:"h"`
	Id          int64 `json:"id"`
	L           HData `json:"l"`
	M           HData `json:"m"`
	Mst         int `json:"mst"`
	Mv          int `json:"mv"`
	Name        string `json:"name"`
	No          int `json:"no"`
	Pop         int `json:"pop"`
	Pst         int `json:"pst"`
	PublishTime int64 `json:"publishTime"`
	Rt          string        `json:"rt"`
	RtUrl       string `json:"rtUrl"`
	RtUrls      []string `json:"rtUrls"`
	Rtype       int `json:"rtype"`
	Rurl        string `json:"rurl"`
	SId         int `json:"s_id"`
	St          int `json:"st"`
	T           int `json:"t"`
	V           int `json:"v"`
}

type HData struct {
	Br   int `json:"br"`
	Fid  int `json:"fid"`
	Size int64 `json:"size"`
	Vd   float64 `json:"vd"`
}

type  ArData struct {
	Alias []string `json:"alias"`
	Id    int `json:"id"`
	Name  string `json:"name"`
	Tns   []string `json:"tns"`
}

type AlData struct {
	Id     int64 `json:"id"`
	Name   string `json:"name"`
	Pic    int64 `json:"pic"`
	PicUrl string `json:"picUrl"`
}

type CreatorData struct {
	AccountStatus      int `json:"accountStatus"`
	AuthStatus         int `json:"authStatus"`
	Authority          int `json:"authority"`
	AvatarImgId        int64 `json:"avatarImgId"`
	AvatarImgIdStr     string `json:"avatarImgIdStr"`
	AvatarImgId_str    string `json:"avatarImgId_str"`
	AvatarUrl          string `json:"avatarUrl"`
	BackgroundImgId    int64 `json:"backgroundImgId"`
	BackgroundImgIdStr string `json:"backgroundImgIdStr"`
	BackgroundUrl      string `json:"backgroundUrl"`
	Birthday           int64 `json:"birthday"`
	City               int `json:"city"`
	DefaultAvatar      bool `json:"defaultAvatar"`
	Description        string `json:"description"`
	DetailDescription  string `json:"detailDescription"`
	DjStatus           int `json:"djStatus"`
	ExpertTags         []string `json:"expertTags"`
	Experts            []string `json:"experts"`
	Followed           bool `json:"followed"`
	Gender             int `json:"gender"`
	Mutual             bool `json:"mutual"`
	Nickname           string `json:"nickname"`
	Province           int `json:"province"`
	RemarkName         []string `json:"remarkName"`
	Signature          string `json:"signature"`
	UserId             int64 `json:"userId"`
	UserType           int `json:"userType"`
}

type TrackIdsData struct {
	Id int64 `json:"id"`
	V  int `json:"v"`
}
// @Title detail
// @Description get user playlist detail
// @Success 200 {object} controllers.PlaylistJson
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
	//beego.Debug(list)
	if len(list) > 0 {
		beego.Debug("使用缓存---")
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
	if len(string(body)) > 0 {
		SetCache(cname, string(body), 600)
	}
	var playlist PlaylistJson
	json.Unmarshal(body,&playlist)
	this.SetReturnData(200, "ok", playlist)
}