package tlog

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Bot struct {
	token string
	admin string
}

type Response struct {
	Ok     bool        `json:"ok"`
	Result interface{} `json:"result"`
}

func (b *Bot) SendMessage(msg string) Response {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=<pre>%s</pre>&parse_mode=html", b.token, b.admin, msg)
	// http request to url
	cli := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	res := Response{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
