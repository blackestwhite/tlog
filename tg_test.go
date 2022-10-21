package tlog

import (
	"net/http"
	"testing"
)

func TestSendMessage(t *testing.T) {
	b := &Bot{
		token:  "5103837914:AAFkNHrB53QWuxl9HX7WJm8H97uOT83RbKQ",
		admin:  "1422816851",
		client: &http.Client{},
	}
	res := b.SendMessage("test")
	if !res.Ok {
		t.Errorf("SendMessage: expected true got %v", res.Ok)
	}
}
