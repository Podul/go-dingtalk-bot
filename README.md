# GO Dingtalk bot

使用

``` golang
    import (
        "dingtalk"
    )

    var bot = dingtalk.NewBot("accessToken", "secret")

    // 发送文本
    res, err := bot.SendText(content string, opts ...BotOption)
    // 发送链接消息
    res, err := bot.SendLink(text string, title string, picUrl string, messageUrl string, opts ...BotOption)

    // 发送markdown消息
    res, err := SendMarkdown(title string, text string, opts ...BotOption)

    // 发送ActionCard消息
    res, err := SendActionCard(title string, text string, singleTitle string, singleURL string, btnOrientation string, opts ...BotOption)

    // 独立跳转ActionCard类型
    res, err := SendActionCard2(title string, text string, btns []ActionCardMessageBtn, btnOrientation string, opts ...BotOption)

    // 发送FeedCard消息
    res, err := SendFeedCard(links []FeedCardMessageLink, opts ...BotOption)
```
