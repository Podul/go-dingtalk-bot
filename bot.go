package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type Bot struct {
	baseURL      string
	access_token string
	secret       string
}

// option
type botOption struct {
	timeout time.Duration

	// 只有文本消息和 markdown 消息能@人
	// @所有人时：true，否则为：false
	isAtAll bool
	// 被@人的手机号
	atMobiles []string
	// 被@人的用户userid
	atUserIds []string
}

type BotOption interface {
	apply(*botOption)
}

type botFuncOption struct {
	f func(*botOption)
}

func (fdo *botFuncOption) apply(do *botOption) {
	fdo.f(do)
}

// 超时时间
func TimeoutOpt(timeout time.Duration) BotOption {
	return &botFuncOption{
		f: func(do *botOption) {
			do.timeout = timeout
		},
	}
}

// @所有人，文本消息和markdown消息支持@所有人
func AtAllOpt(isAtAll bool) BotOption {
	return &botFuncOption{
		f: func(do *botOption) {
			do.isAtAll = isAtAll
		},
	}
}

// @手机号，文本消息和markdown消息支持@手机号
func AtMobilesOpt(atMobiles []string) BotOption {
	return &botFuncOption{
		f: func(do *botOption) {
			do.atMobiles = atMobiles
		},
	}
}

// @用户id，文本消息和markdown消息支持@用户
func AtUserIdsOpt(atUserIds []string) BotOption {
	return &botFuncOption{
		f: func(do *botOption) {
			do.atUserIds = atUserIds
		},
	}
}

type BotRes = map[string]interface{}

func NewBot(access_token string, secret string) *Bot {
	return &Bot{access_token: access_token, secret: secret, baseURL: "https://oapi.dingtalk.com/robot/send"}
}

// 发送文本消息
func (bot *Bot) SendText(content string, opts ...BotOption) (BotRes, error) {
	return bot.send(TextMessage{Content: content}, opts...)
}

// 发送链接消息
func (bot *Bot) SendLink(text string, title string, picUrl string, messageUrl string, opts ...BotOption) (BotRes, error) {
	return bot.send(LinkMessage{Text: text, Title: title, PicUrl: picUrl, MessageUrl: messageUrl}, opts...)
}

// 发送markdown消息
func (bot *Bot) SendMarkdown(title string, text string, opts ...BotOption) (BotRes, error) {
	return bot.send(MarkdownMessage{Title: title, Text: text}, opts...)
}

// 发送ActionCard消息
func (bot *Bot) SendActionCard(title string, text string, singleTitle string, singleURL string, btnOrientation string, opts ...BotOption) (BotRes, error) {
	return bot.send(ActionCardMessage{ActionCard: ActionCard{Title: title, Text: text, BtnOrientation: btnOrientation}, SingleTitle: singleTitle, SingleURL: singleURL}, opts...)
}

// 独立跳转ActionCard类型
func (bot *Bot) SendActionCard2(title string, text string, btns []ActionCardMessageBtn, btnOrientation string, opts ...BotOption) (BotRes, error) {
	return bot.send(ActionCardMessage2{ActionCard: ActionCard{Title: title, Text: text, BtnOrientation: btnOrientation}, Btns: btns}, opts...)
}

// 发送FeedCard消息
func (bot *Bot) SendFeedCard(links []FeedCardMessageLink, opts ...BotOption) (BotRes, error) {
	return bot.send(FeedCardMessage{Links: links}, opts...)
}

// 发送消息
func (bot *Bot) send(message any, opts ...BotOption) (BotRes, error) {
	var msgType string
	switch m := message.(type) {
	case TextMessage:
		msgType = msgTypeText
	case string:
		msgType = msgTypeText
		message = TextMessage{Content: m}
	case LinkMessage:
		msgType = msgTypeLink
	case MarkdownMessage:
		msgType = msgTypeMarkdown
	case ActionCardMessage:
		msgType = msgTypeActionCard
	case ActionCardMessage2:
		msgType = msgTypeActionCard
	case FeedCardMessage:
		msgType = msgTypeFeedCard
	default:
		return nil, errors.New("message type error")
	}

	options := botOption{timeout: 15 * time.Second, isAtAll: false, atMobiles: []string{}, atUserIds: []string{}}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt.apply(&options)
	}

	sign, timestamp := bot.genSign()
	params := map[string]string{
		"access_token": bot.access_token,
		"timestamp":    timestamp,
		"sign":         sign,
	}

	body := map[string]any{
		"msgtype": msgType,
		msgType:   anyToMap(message),
	}

	if options.isAtAll || len(options.atMobiles) > 0 || len(options.atUserIds) > 0 {
		body["at"] = map[string]interface{}{
			"isAtAll":   options.isAtAll,
			"atMobiles": options.atMobiles,
			"atUserIds": options.atUserIds,
		}
	}

	client := newClient(ClientConfig{Timeout: options.timeout})
	response := BotRes{}
	err := client.request("POST", bot.baseURL, params, body, &response)
	return response, err
}

// 签名
func (bot *Bot) genSign() (string, string) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	data := []byte(timestamp + "\n" + bot.secret)
	mac := hmac.New(sha256.New, []byte(bot.secret))
	_, _ = mac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), timestamp
}

func anyToMap(data any) map[string]any {
	b, _ := json.Marshal(data)
	m := map[string]any{}
	_ = json.Unmarshal(b, &m)
	return m
}
