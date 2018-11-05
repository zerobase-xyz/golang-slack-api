package main

type Params struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

type Slack struct {
	params Params
	url    string
}

type Reslack struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
