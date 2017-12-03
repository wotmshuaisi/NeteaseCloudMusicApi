package controllers

import (
	"github.com/ActingCute/NeteaseCloudMusicApi/models"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	baseController
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

//form.Set("offset", "0")
//form.Set("uid", this.GetString("uid"))
//form.Set("limit", "1000")
//form.Set("csrf_token", "")

type Playlist struct {
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
	Experts            int`json:"experts"`        //{}
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

const clientToken = "1_jVUMqWEPke0/1/Vu56xCmJpo5vP1grjn_SOVVDzOc78w8OKLVZ2JH7IfkjSXqgfmh"

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title CellphoneLogin
// @Description Logs user into the system
// @Param	phone			query 	string	true		"The phone    for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 400 param error
// @Failure 500 Program error
// @Failure 501 Account number does not exist
// @Failure 502 Password error
// @router /cellphonelogin [get]
func (this *UserController) CellphoneLogin() {
	if this.IsLogin() {
		this.SetReturnData(200, "login success", UserInfo)
		return
	}
	this.NeedAPIinput("phone", "password")
	var loginjson phonelogin
	loginjson.Password = Md5(this.GetString("password"))
	loginjson.Phone = this.GetString("phone")
	loginjson.RememberLogin = true
	loginjson.ClientToken = clientToken
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
	beego.Debug("Cookies===", this.Ctx.Request.Cookies())

	//for _,c := range Cookies {
	//	this.Ctx.Request.AddCookie(c)
	//}

	this.SetReturnData(200, "login success", UserInfo)
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title playlist
// @Description get user playlist
// @Success 200 {string} get playlist success
// @router /playlist [get]
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

	var playlist Playlist
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
	SetCache(cname,string(body),600)
	this.SetReturnData(200, "ok", string(body))
}