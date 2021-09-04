package ncm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/secriy/ncmlu-api/util"
)

type NCMAccount struct {
	Phone            string `json:"phone"`       // 用户手机号
	Password         string `json:"password"`    // 用户密码
	CountryCode      int    `json:"countrycode"` // 用户手机号国家码
	Csrf             string // CSRF Token
	Nickname         string // 昵称
	Uid              int    // 用户 UID
	Level            int    // 用户等级
	RemainPlayCount  int    // 剩余播放数量
	RemainLoginCount int    // 剩余登录天数
	PlayList         []int  // 歌单列表
	MusicList        []int  // 歌曲列表
}

// New return a NCMAccount instance
func New(phone, password string, code int) *NCMAccount {
	return &NCMAccount{
		Phone:       phone,
		Password:    password,
		CountryCode: code,
	}
}

// Login 登录
func (ac *NCMAccount) Login(client *http.Client) {
	loginURL := "https://music.163.com/weapi/login/cellphone"
	jsonData, err := json.Marshal(ac)
	if err != nil {
		util.Logger.Error(err)
		return
	}
	req, err := postReq(loginURL, string(jsonData))
	if err != nil {
		util.Logger.Error(err)
		return
	}
	req.Header.Set("Cookie", "os=pc; osver=Microsoft-Windows-10-Professional-build-10586-64bit; appver=2.0.3.131777; channel=netease; __remember_me=true;")
	res, err := client.Do(req)
	if err != nil {
		util.Logger.Error(err)
		return
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
		util.Logger.Error(err)
		return
	}
	// serializer
	data := ResponseJSON{}
	if err := json.Unmarshal(body, &data); err != nil {
		util.Logger.Error(err)
		return
	}
	if data.Account.ID == 0 {
		util.Logger.Errorf("用户[%s]登录失败: %s", ac.Phone, string(body))
		return
	}
	// set id and name
	ac.Uid, ac.Nickname = data.Account.ID, data.Profile.Nickname
	ac.UserLevel(client)

	util.Logger.Infof("用户[%s][%s]登录成功，当前等级 %d", ac.Phone, ac.Nickname, ac.Level)

	util.Logger.Infof("用户[%s][%s]剩余登录天数：%d，剩余听歌数：%d", ac.Phone, ac.Nickname, ac.RemainLoginCount, ac.RemainPlayCount)
}

// Sign 签到
func (ac *NCMAccount) Sign(client *http.Client, tp int) {
	signURL := "https://music.163.com/weapi/point/dailyTask?" + ac.Csrf
	// 签到类型
	t := "安卓端"
	if tp == 1 {
		t = "PC/Web 端"
	}
	// request
	res, err := postRes(signURL, fmt.Sprintf("{\"type\": %d}", tp), client)
	if err != nil {
		util.Logger.Errorf("用户[%s][%s]%s签到失败: %s", ac.Phone, ac.Nickname, t, err)
		return
	}
	// serializer
	data := ResponseJSON{}
	if err = json.Unmarshal(res, &data); err != nil {
		util.Logger.Errorf("用户[%s][%s]%s签到失败: %s", ac.Phone, ac.Nickname, t, err)
		return
	}
	// results
	if data.Code == 200 {
		util.Logger.Infof("用户[%s][%s]%s签到成功", ac.Phone, ac.Nickname, t)
		return
	} else if data.Code == -2 {
		util.Logger.Infof("用户[%s][%s]%s重复签到", ac.Phone, ac.Nickname, t)
		return
	}
	util.Logger.Errorf("用户[%s][%s]%s签到失败: %s", ac.Phone, ac.Nickname, t, string(res))
}

// UserLevel 用户等级
func (ac *NCMAccount) UserLevel(client *http.Client) {
	levelURL := "https://music.163.com/weapi/user/level?csrf_token=" + ac.Csrf
	res, err := postRes(levelURL, "{}", client)
	if err != nil {
		util.Logger.Error(err)
		return
	}
	// serializer
	data := ResLevelJSON{}
	if err = json.Unmarshal(res, &data); err != nil {
		util.Logger.Error(err)
		return
	}
	ac.Level = data.Data.Level
	ac.RemainPlayCount = data.Data.NextPlayCount - data.Data.NowPlayCount
	ac.RemainLoginCount = data.Data.NextLoginCount - data.Data.NowLoginCount
}

// RecommendList 获取推荐歌单列表
func (ac *NCMAccount) RecommendList(client *http.Client) {
	recommendURL := "https://music.163.com/weapi/discovery/recommend/resource?csrf_token=" + ac.Csrf
	res, err := postRes(recommendURL, fmt.Sprintf(`{"csrf_token":"%s"}`, ac.Csrf), client)
	if err != nil {
		util.Logger.Error(err)
		return
	}
	data := RecommendJSON{}
	if err := json.Unmarshal(res, &data); err != nil {
		util.Logger.Error(err)
		return
	}
	playList := make([]int, len(data.Recommend))
	for k, v := range data.Recommend {
		playList[k] = v.ID
	}
	ac.PlayList = append(ac.PlayList, playList...)
}

// Musics 获取某一歌单中的所有歌曲
func (ac *NCMAccount) Musics(client *http.Client, listID int) {
	detailURL := "https://music.163.com/weapi/v6/playlist/detail?csrf_token=" + ac.Csrf
	res, err := postRes(detailURL, fmt.Sprintf(`{"id": %d,"n": 1000,"csrf_token": "%s"}`, listID, ac.Csrf), client)
	if err != nil {
		util.Logger.Error(err)
		return
	}
	data := MusicListJSON{}
	if err := json.Unmarshal(res, &data); err != nil {
		util.Logger.Error(err)
		return
	}
	musicList := make([]int, len(data.PlayList.TrackIDs))
	for k, v := range data.PlayList.TrackIDs {
		musicList[k] = v.ID
	}
	ac.MusicList = append(ac.MusicList, musicList...)
}

func (ac *NCMAccount) Feedback(client *http.Client) {
	feedbackURL := "http://music.163.com/weapi/feedback/weblog"
	logs := make([]Log, len(ac.MusicList))
	for k, v := range ac.MusicList {
		l := Log{
			Action: "play",
			Json:   Json{Download: 0, End: "playend", ID: v, SourceID: "", Time: 240, Type: "song", WiFi: 0},
		}
		logs[k] = l
	}
	data, err := json.Marshal(logs)
	if err != nil {
		util.Logger.Error(err)
		return
	}
	res, err := postRes(feedbackURL, fmt.Sprintf("{'logs':'%s'}", string(data)), client)
	if err != nil {
		util.Logger.Error(err)
		return
	}
	jsonData := FeedbackJSON{}
	if err := json.Unmarshal(res, &jsonData); err != nil {
		util.Logger.Error(err)
		return
	}
	if jsonData.Code == 200 {
		util.Logger.Infof("用户[%s][%s]刷听歌量成功", ac.Phone, ac.Nickname)
	} else {
		util.Logger.Errorf("用户[%s][%s]刷听歌量失败: %s", ac.Phone, ac.Nickname, string(res))
	}
}
