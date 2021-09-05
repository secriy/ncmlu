package ncm

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

type RecommendJSON struct {
	Code      int `json:"code"`
	Recommend []struct {
		ID int `json:"id"`
	} `json:"recommend"`
}

type PersonalizedJSON struct {
	Code   int `json:"code"`
	Result []struct {
		ID int `json:"id"`
	} `json:"result"`
}

type MusicListJSON struct {
	PlayList struct {
		TrackIDs []struct {
			ID int `json:"id"`
		} `json:"trackIds"`
	} `json:"playlist"`
}

type FeedbackJSON struct {
	Code int `json:"code"`
}
