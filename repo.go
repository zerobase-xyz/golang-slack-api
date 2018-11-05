package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func NewSlack(url string, text string, username string, iconEmoji string, iconURL string, channel string) *Slack {

	p := Params{
		Text:      text,
		Username:  username,
		IconEmoji: iconEmoji,
		IconURL:   iconURL,
		Channel:   channel}

	return &Slack{
		params: p,
		url:    url}
}

func (s *Slack) Send() {

	params, _ := json.Marshal(s.params)

	resp, err := http.PostForm(
		s.url,
		url.Values{"payload": {string(params)}},
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
}
