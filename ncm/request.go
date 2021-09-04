package ncm

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/secriy/ncmlu-api/util"
	"golang.org/x/net/publicsuffix"
)

type Log struct {
	Action string `json:"action"`
	Json   Json   `json:"json"`
}

type Json struct {
	Download int    `json:"download"`
	End      string `json:"end"`
	ID       int    `json:"id"`
	SourceID string `json:"sourceId"`
	Time     int    `json:"time"`
	Type     string `json:"type"`
	WiFi     int    `json:"wifi"`
}

func NewClient() *http.Client {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client := &http.Client{Jar: jar}
	return client
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

func postRes(_url, postForm string, client *http.Client) ([]byte, error) {
	req, err := postReq(_url, postForm)
	if err != nil {
		util.Logger.Error("post error: " + err.Error())
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		util.Logger.Error("post error: " + err.Error())
		return nil, err
	}
	defer res.Body.Close()
	// read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		util.Logger.Error("post error: " + err.Error())
		return nil, err
	}
	return body, nil
}
