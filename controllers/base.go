package controllers

import (
	"github.com/astaxie/beego"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/gob"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"github.com/astaxie/beego/i18n"
	"github.com/ActingCute/NeteaseCloudMusicApi/models"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

type RestfulReturn struct {
	Result  int16
	Message string
	Data    interface{}
}

type langType struct {
	Lang, Name string
}

var (
	BaseApi = "http://music.163.com"
	PhoneLoginApi = "/weapi/login/cellphone"
	langTypes   []*langType
)

func init() {
	initLang()
}

//语言配置
func initLang() {

	langs := strings.Split(models.GetAppConf("lang::types"), "|")
	names := strings.Split(models.GetAppConf("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))

	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		if err := i18n.SetMessage(lang, beego.AppPath + "\\languages\\default\\" + lang + ".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		} else {
			beego.Trace("Loading language" + lang)
		}
	}

}

func (this *baseController)HttpPost(url string, data *strings.Reader) ([]byte, error) {
	request, err := http.NewRequest("POST", url, data)
	if err != nil {
		beego.Error("Post NewRequest ,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}

	request = this.getRequestHeader(request)

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)

	if err != nil {
		beego.Error("Post http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("Post ReadAll failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	cookie := resp.Header["Set-Cookie"]
	if len(cookie) > 0{
		for _, c := range resp.Cookies() {
			request.AddCookie(c)
		}
	}
	return b, err
}

func (this *baseController)getRequestHeader(request *http.Request) *http.Request {
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
	request.Header.Set("Connection", "Keep-Alive")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Host", "music.163.com")
	request.Header.Set("Cookie", "cookie")
	request.Header.Set("User-Agent", this.RandomUserAgent())
	return request
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (this *baseController)Md5(str string) string {
	if len(str) < 1 {
		//什么都没有,不给你加密
		return str
	}
	c := md5.New()
	c.Write([]byte(str)) // 需要加密的字符串
	cipherStr := c.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//随机客户端
func (this *baseController)RandomUserAgent() string {
	userAgentList := []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
		"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89; GameHelper",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
		"Mozilla/5.0 (Windows NT 6.3; Win64, x64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
		"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1"}
	num := rand.Intn(len(userAgentList))
	return userAgentList[num]
}

func (this *baseController) NeedAPIinput(fields ...string) {

	for _, field := range fields {

		if len(this.GetString(field)) == 0 {
			resultStringList := map[string]string{
				"phone":  "please_enter_phone",
				"password":  "please_enter_password",
			}
			beego.Debug("缺少参数-----", resultStringList[field])
			this.SetReturnData(400, this.L(resultStringList[field]), nil)
			this.StopRun()
		}
	}

}

func (this *baseController) SetReturnData(result int, message string, data interface{}) {

	rt := &RestfulReturn{Result: int16(result), Message: message, Data: data}
	this.Data["json"] = rt
	this.ServeJSON()

}
func (this *baseController) L(name string) string {

	return this.Tr(name)

}

//处理返回code
func (this *baseController)StatusCode(code int) bool {
	switch code {
	case 200:
		return true
	case 502:
		return false
	default:
		return false
	}
}

//判断登录是否过期
//func (this *baseController)IsLogin() bool{
//	if UserInfo.Account.Id < 1 {
//		return false
//	}
//	t := UserInfo.Account.Type
//	for _,v := range UserInfo.Bindings {
//		if v.Type == t {
//
//		}
//	}
//}