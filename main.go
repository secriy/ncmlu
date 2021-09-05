package main

import (
	"time"

	"github.com/secriy/ncmlu/config"
	"github.com/secriy/ncmlu/ncm"
	"github.com/secriy/ncmlu/util"
)

func main() {
	config.InitConfig()
	util.InitLogger("info")

	for k, v := range config.Conf.Accounts {
		execTask(v.Phone, v.Passwd, v.Expired, !v.OnlySign)
		if k > 0 && k%20 == 0 {
			time.Sleep(time.Minute * 5)
		}
	}
}

func execTask(phone, passwd, expired string, play bool) {
	t, err := time.Parse("2006-01-02", expired)
	if err != nil {
		util.Logger.Errorf("%s expired time parsing error: %s", phone, err)
		return
	}
	if t.Before(time.Now()) {
		return
	}
	ncm.NcmluTask(phone, passwd, 86, play)
}
