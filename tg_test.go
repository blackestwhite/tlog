package tlog

import "testing"

func TestSendMessage(t *testing.T) {
	wg.Add(1)
	b := &Bot{
		token: "5103837914:AAFkNHrB53QWuxl9HX7WJm8H97uOT83RbKQ",
		admin: "1422816851",
	}
	res := b.SendMessage("test")
	if !res.Ok {
		t.Errorf("SendMessage: expected true got %v", res.Ok)
	}
}
