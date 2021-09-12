package main

import (
	"time"

	"github.com/secriy/ncmlu/config"
	"github.com/secriy/ncmlu/ncm"
	"github.com/secriy/ncmlu/util"
)

func main() {
	config.InitConfig()
	util.InitLogger(config.Conf.Level)

	for k, v := range config.Conf.Accounts {
		execTask(v.Phone, v.Passwd, v.Expired, v.OnlySign, v.Unstable)
		time.Sleep(time.Second * 2)
		if k > 0 && k%25 == 0 {
			// sleep 2 minute every 25 accounts
			time.Sleep(time.Minute * 2)
		}
	}
}

func execTask(phone, passwd, expired string, onlySign, unstable bool) {
	t, err := time.Parse("2006-01-02", expired)
	if err != nil {
		util.Logger.Errorf("%s expired time parsing error: %s", phone, err)
		return
	}
	if t.Before(time.Now()) {
		return
	}
	ncm.NcmluTask(phone, passwd, 86, onlySign, unstable)
}
