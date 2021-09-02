package ncmlu_api

import (
	"bytes"
	"encoding/json"

	"github.com/secriy/ncmlu-api/util"
)

func NcmluTask(phone, passwd string, code int) {
	util.InitLogger("info")

	client := NewClient()

	acc := New(phone, util.MD5Sum(passwd), code)
	// 登录
	acc.Login(client)
	if acc.Uid == 0 {
		return
	}
	// 签到
	acc.Sign(client, 0) // 安卓端
	acc.Sign(client, 1) // PC/Web 端
	// 获取推荐歌单
	acc.RecommendList(client)
	if acc.PlayList == nil || len(acc.PlayList) == 0 {
		return
	}
	// 获取全部歌曲
	for _, v := range acc.PlayList {
		acc.Musics(client, v)
	}
	if acc.MusicList == nil || len(acc.MusicList) == 0 {
		return
	}
	if len(acc.MusicList) > 310 {
		acc.MusicList = acc.MusicList[:310]
	}
	// 刷歌
	acc.Feedback(client)

	bs, _ := json.Marshal(acc)
	var out bytes.Buffer
	_ = json.Indent(&out, bs, "", "\t")
	util.Logger.Debug(out.String())
}
