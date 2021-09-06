package ncm

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/secriy/ncmlu/config"
	"github.com/secriy/ncmlu/util"
)

const limit = 410

func NcmluTask(phone, passwd string, code int, play bool) {
	util.InitLogger("info")

	client := NewClient()

	if len(passwd) != 32 {
		passwd = util.MD5Sum(passwd)
	}

	acc := New(phone, passwd, code)
	// 登录
	acc.Login(client)
	if acc.Uid == 0 {
		return
	}
	// 签到
	acc.Sign(client, 0) // 安卓端
	acc.Sign(client, 1) // PC/Web 端

	if play {
		// 获取歌单
		if config.Conf.Playlist == nil || len(config.Conf.Playlist) == 0 {
			// 获取个性推荐推荐歌单
			acc.PersonalizedList(client)
			// 获取全部歌曲
			for _, v := range acc.PlayList {
				acc.Musics(client, v)
			}
			if acc.MusicList == nil || len(acc.MusicList) == 0 {
				return
			}
			// 限制歌曲数量
			if len(acc.MusicList) > limit {
				randomMusics(acc.MusicList, limit)
				acc.MusicList = acc.MusicList[:limit]
			}
		} else {
			// 获取自定义歌单
			acc.PlayList = config.Conf.Playlist
			// 获取全部歌曲
			for _, v := range acc.PlayList {
				acc.Musics(client, v)
			}
			if acc.MusicList == nil || len(acc.MusicList) == 0 {
				return
			}
		}
		// 刷歌
		acc.Feedback(client)
	}

	bs, _ := json.Marshal(acc)
	var out bytes.Buffer
	_ = json.Indent(&out, bs, "", "\t")
	util.Logger.Debug(out.String())
}

// randomMusics reshuffle the music slice.
func randomMusics(musics []int, num int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		ri := rand.Intn(len(musics))
		musics[i], musics[ri] = musics[ri], musics[i]
	}
}
