# tlog
golang log but via telegram bot support

## how to use
```go
tlog.LinkBot("BOT TOKEN", "CHAT ID", &http.Client{})
tlog.Print("Hello tlog")
```
**CHAT ID** is your telegram user identifier, you can find yours on [@userinfobot](https://t.me/userinfobot)

## custom logger
```go
var b bytes.Buffer
l := tlog.New(&b, "prefix:", LstdFlags)
l.LinkBot("BOT TOKEN", "CHAT ID", &http.Client{})
```