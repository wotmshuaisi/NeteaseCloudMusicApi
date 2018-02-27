package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
)

// Operations about User
type UserController struct {
	BaseController
}

type login struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	RememberLogin bool `json:"rememberLogin"`
	ClientToken   string `json:"clientToken"`
}

type phonelogin struct {
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	RememberLogin bool `json:"rememberLogin"`
	ClientToken   string `json:"clientToken"`
}

type UserData struct {
	Uid        int64 `json:"uid"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Csrf_token string `json:"csrf_token"`
}

type User struct {
	LoginType int `json:"loginType"`
	Code      int `json:"code"`
	Account   UserAccount
	Profile   UserProfile
	Bindings  []UserBindings
	Msg       string `json:"msg"`
}

type UserAccount struct {
	Id                 int64 `json:"id"`
	UserName           string`json:"userName"`
	Type               int `json:"type"`
	Status             int `json:"status"`
	WhitelistAuthority int `json:"whitelistAuthority"`
	CreateTime         int64 `json:"createTime"`
	Salt               string `json:"salt"`
	TokenVersion       int `json:"tokenVersion"`
	Ban                int `json:"ban"`
	BaoyueVersion      int `json:"baoyueVersion"`
	DonateVersion      int`json:"donateVersion"`
	VipType            int `json:"vipType"`
	ViptypeVersion     int `json:"viptypeVersion"`
	AnonimousUser      bool `json:"anonimousUser"`
}

type UserProfile struct {
	DetailDescription  string `json:"detailDescription"`
	DjStatus           int `json:"djStatus"`
	Followed           bool `json:"followed"`
	Description        string `json:"description"`
	AvatarImgId        int64 `json:"avatarImgId"`
	ExpertTags         string `json:"expertTags"` //null
	AuthStatus         int `json:"authStatus"`
	BackgroundImgId    int64 `json:"backgroundImgId"`
	UserType           int `json:"userType"`
	Experts            string `json:"experts"`        //{}
	BackgroundUrl      string `json:"backgroundUrl"`
	AvatarImgIdStr     int64 `json:"avatarImgIdStr"`
	BackgroundImgIdStr int64 `json:"backgroundImgIdStr"`
	UserId             int64 `json:"userId"`
	AccountStatus      int `json:"accountStatus"`
	Nickname           string `json:"nickname"`
	RemarkName         string `json:"remarkName"` //null
	Mutual             bool `json:"mutual"`
	Province           int64 `json:"province"`
	DefaultAvatar      bool `json:"defaultAvatar"`
	AvatarUrl          string `json:"avatarUrl"`
	Gende              int `json:"gende"`
	Birthday           int64 `json:"birthday"`
	City               int64 `json:"city"`
	VipType            int `json:"vipType"`
	Signature          string`json:"signature"`
	Authority          int `json:"authority"`
	AvatarImgId_str    int64 `json:"avatarImgId_str"`
}

type UserBindings struct {
	ExpiresIn    int64 `json:"expiresIn"`
	RefreshTime  int64 `json:"refreshTime"`
	Url          string `json:"url"`
	UserId       int64 `json:"userId"`
	TokenJsonStr UserTokenJsonStr
	Expired      bool `json:"expired"`
	Id           int64 `json:"id"`
	Type         int `json:"type"`
}

type UserTokenJsonStr struct {
	Countrycode string `json:"countrycode"`
	Cellphone   int64 `json:"cellphone"`
	HasPassword bool `json:"hasPassword"`
}

var (
	UserInfo User
)

// @Title CellphoneLogin
// @Description 用手机号码登陆
// @Param	phone			query 	string	true		"The phone    for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {object} controllers.User
// @Failure 400 param error
// @Failure 500 Program error
// @Failure 501 Account number does not exist
// @Failure 502 Password error
// @router /cellphonelogin [get]
func (this *UserController) CellphoneLogin() {
	if this.IsLogin() {
		beego.Debug("be login")
		this.SetReturnData(200, "login success", UserInfo)
		return
	}
	this.NeedAPIinput("phone", "password")
	var loginjson phonelogin
	loginjson.Password = Md5(this.GetString("password"))
	loginjson.Phone = this.GetString("phone")
	loginjson.RememberLogin = true
	loginjson.ClientToken = ClientToken
	logindata, err := json.Marshal(loginjson)
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	body, err := this.Http(BaseApi + PhoneLoginApi, logindata, "POST")
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	json.Unmarshal(body, &UserInfo)
	if !this.StatusCode(UserInfo.Code) {
		msg := UserInfo.Msg
		if UserInfo.Code == 400 {
			msg = "Account_or_password_error"
		}
		this.SetReturnData(UserInfo.Code, msg, nil)
		return
	}
	this.Ctx.SetCookie("uid", strconv.FormatInt(UserInfo.Account.Id, 10), -1, "/")
	if !this.IsLogin() {
		this.SetReturnData(200, "login fail", this.Ctx.Request.Cookies())
	}
	beego.Debug("Cookies===", Cookies)
	this.SetReturnData(200, "login success", UserInfo)
}

// @Title playlist
// @Description 获取用户歌单
// @Success 200 {string} get playlist success
// @router /playlist [get]
// @Param uid query string true "The uid for user playlist"
func (this *UserController) Playlist() {
	this.NeedAPIinput("uid")
	uid, _ := this.GetInt64("uid")
	if uid < 1 {
		this.SetReturnData(500, "uid error", nil)
		return
	}
	uidstr := this.GetString("uid")
	cname := uidstr + "playlist"
	list := GetCache(cname)
	//beego.Debug(list)
	if len(list) > 0 {
		this.SetReturnData(200, "ok", list)
		return
	}

	var playlist UserData
	playlist.Uid = uid
	playlist.Csrf_token = ""
	playlist.Limit = 1000
	playlist.Offset = 0

	playlistdata, err := json.Marshal(playlist)
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}

	body, err := this.Http(BaseApi + MusicListApi, playlistdata, "POST")
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	SetCache(cname, body, 600)
	this.SetReturnData(200, "ok", string(body))
}

// @Title detail
// @Description 获取用户详情
// @Success 200 {string} get user detail success
// @router /detail [get]
// @Param uid query string true "The uid for user detail"
func (this *UserController)Detail() {
	this.NeedAPIinput("uid")
	uid, _ := this.GetInt64("uid")
	if uid < 1 {
		this.SetReturnData(500, "uid error", nil)
		return
	}
	uidstr := this.GetString("uid")
	cname := uidstr + "detail"
	list := GetCache(cname)
	beego.Debug(list)
	usecache ,_:= this.GetInt64("refresh")
	if len(list) > 0 && usecache == 0 {
		this.SetReturnData(200, "ok", string(list))
		return
	}
	var detail UserData
	detail.Csrf_token = ""
	detaildata, err := json.Marshal(detail)
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	body, err := this.Http(BaseApi + V1Dir + UserDetailApi + uidstr, detaildata, "POST")
	if err != nil {
		this.SetReturnData(500, "Program error", err)
		return
	}
	SetCache(cname, body, 600)
	this.SetReturnData(200, "ok", string(body))
}