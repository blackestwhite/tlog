package tlog

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type Bot struct {
	token string
	admin string
}

func (b *Bot) SendMessage(msg string) {
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
	wg.Done()
	defer resp.Body.Close()
}
