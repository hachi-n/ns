package pigeon

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/pigeon/lib/api"
	"io/ioutil"
)

type Slack struct {
	message   string
	url       string
	channel   string
	messenger string
	icon      SlackIcon
}

type SlackIcon struct {
	emoji string
	url   string
}

func NewSlack(message, url, channel, messenger, iconEmoji, iconUrl string) *Slack {
	return &Slack{
		message:   message,
		url:       url,
		channel:   channel,
		messenger: messenger,
		icon:      SlackIcon{emoji: iconEmoji, url: iconUrl},
	}
}

func (s *Slack) MarshalJSON() ([]byte, error) {
	marshalStruct := &struct {
		Channel   string `json:"channel"`
		Username  string `json:"username"`
		Text      string `json:"text"`
		IconEmoji string `json:"icon_emoji,omitempty"`
		IconUrl   string `json:"icon_url,omitempty"`
	}{
		s.channel,
		s.messenger,
		s.message,
		s.icon.emoji,
		s.icon.url,
	}

	return json.Marshal(marshalStruct)
}

func (s *Slack) Notify() error {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	payload, err := json.Marshal(s)
	if err != nil {
		return err
	}

	resp, err := api.Request(s.url, payload, api.Post, headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
