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

	catnap := config.Conf.Catnap
	sleep := config.Conf.Sleep

	length := len(config.Conf.Accounts)

	for k, v := range config.Conf.Accounts {
		execTask(v.Phone, v.Passwd, v.Code, v.Expired, v.OnlySign, v.Unstable)

		if k == length-1 {
			// break after execute last one
			break
		}
		if k > 0 && sleep.Number > 0 && sleep.Duration > 0 && k%sleep.Number == 0 {
			// sleep
			time.Sleep(time.Minute * time.Duration(sleep.Duration))
		} else if k > 0 && catnap.Number > 0 && catnap.Duration > 0 && k%catnap.Number == 0 {
			// catnap
			time.Sleep(time.Minute * time.Duration(catnap.Duration))
		} else {
			time.Sleep(time.Second * time.Duration(config.Conf.Interval))
		}
	}
}

func execTask(phone, passwd string, code int, expired string, onlySign, unstable bool) {
	t, err := time.Parse("2006-01-02", expired)
	if err != nil {
		util.Logger.Errorf("%s expired time parsing error: %s", phone, err)
		return
	}
	if t.Before(time.Now()) {
		return
	}
	if code == 0 {
		// set default country code
		code = 86
	}
	ncm.NcmluTask(phone, passwd, code, onlySign, unstable)
}
