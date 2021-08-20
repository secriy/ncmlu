package ncmlu_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/secriy/ncmlu-api/util"
)

type NCMAccount struct {
	Phone            string `json:"phone"`
	Password         string `json:"password"`
	CountryCode      int    `json:"countrycode"`
	Csrf             string
	Nickname         string
	Uid              int
	Level            int
	RemainPlayCount  int
	RemainLoginCount int
	SignRet          string
}

type ResponseJSON struct {
	Code    int `json:"code"`
	Account struct {
		ID int `json:"id"`
	} `json:"account"`
	Profile struct {
		Nickname string `json:"nickname"`
	} `json:"profile"`
}

type ResLevelJSON struct {
	Code int `json:"code"`
	Data struct {
		Level          int `json:"level"`
		NextLoginCount int `json:"nextLoginCount"`
		NextPlayCount  int `json:"nextPlayCount"`
		NowLoginCount  int `json:"nowLoginCount"`
		NowPlayCount   int `json:"nowPlayCount"`
	} `json:"data"`
}

// New return a NCMAccount instance
func New(phone, password string, code int) *NCMAccount {
	return &NCMAccount{
		Phone:       phone,
		Password:    password,
		CountryCode: code,
	}
}

// postReq do post job
func postReq(_url, postForm string) (*http.Request, error) {
	params, encSecKey, err := util.EncryptForm(postForm)
	if err != nil {
		return nil, err
	}
	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	payload := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", _url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")
	req.Header.Add("Referer", "http://music.163.com/")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	return req, nil
}

// Login 登录
func (ac *NCMAccount) Login(client *http.Client) error {
	loginURL := "https://music.163.com/weapi/login/cellphone"
	jsonData, err := json.Marshal(ac)
	if err != nil {
		return err
	}
	req, err := postReq(loginURL, string(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", "os=pc; osver=Microsoft-Windows-10-Professional-build-10586-64bit; appver=2.0.3.131777; channel=netease; __remember_me=true;")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// set csrf token
	for _, v := range res.Cookies() {
		if v.Name == "__csrf" {
			ac.Csrf = v.Value
		}
	}
	// read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// serializer
	data := ResponseJSON{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	// set id and name
	ac.Uid, ac.Nickname = data.Account.ID, data.Profile.Nickname
	err = ac.UserLevel(client)
	return err
}

// Sign 签到
func (ac *NCMAccount) Sign(client *http.Client, tp int) (err error) {
	signURL := "https://music.163.com/weapi/point/dailyTask?" + ac.Csrf
	req, err := postReq(signURL, fmt.Sprintf("{\"type\": %d}", tp))
	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	// read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	// serializer
	data := ResponseJSON{}
	if err = json.Unmarshal(body, &data); err != nil {
		return
	}
	// results
	if data.Code == 200 {
		if tp == 0 {
			ac.SignRet = "安卓端签到成功"
		} else {
			ac.SignRet = "PC/Web端签到成功"
		}
	} else if data.Code == -2 {
		ac.SignRet = "重复签到"
	}
	return
}

// UserLevel return the account level.
func (ac *NCMAccount) UserLevel(client *http.Client) error {
	levelURL := "https://music.163.com/weapi/user/level?csrf_token=" + ac.Csrf
	req, err := postReq(levelURL, "{}")
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// serializer
	data := ResLevelJSON{}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println("dsf" + err.Error())
		return err
	}
	ac.Level = data.Data.Level
	ac.RemainPlayCount = data.Data.NextPlayCount - data.Data.NowPlayCount
	ac.RemainLoginCount = data.Data.NextLoginCount - data.Data.NowLoginCount
	return nil
}
