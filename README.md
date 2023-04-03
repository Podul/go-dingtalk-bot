# GO Dingtalk bot

1. 使用
```
	bot := dingtalk.NewBot("accessToken", "secret")
    res, err := bot.SendText("text message", dingtalk.AtAllOpt(true))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(res)
```
