package pigeon

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/pigeon/lib/api"
	"io/ioutil"
)

type Slack struct {
	url       string
	message   string
	channel   string
	messenger string
	icon      SlackIcon
}

type SlackIcon struct {
	emoji string
	url   string
}

func NewSlack(url, message, channel, messenger, iconEmoji, iconUrl string) *Slack {
	return &Slack{
		url:       url,
		message:   message,
		channel:   channel,
		messenger: messenger,
		icon:      SlackIcon{emoji: iconEmoji, url: iconUrl},
	}
}

func (s *Slack) MarshalJSON() ([]byte, error) {
	marshalStruct := struct {
		channel   string
		username  string
		text      string
		iconEmoji string
		iconUrl   string
	}{s.channel,
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

	fmt.Println(body)
	return nil
}
