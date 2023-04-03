package dingtalk

const (
	// Text 类型消息
	msgTypeText = "text"
	// Link 类型消息
	msgTypeLink = "link"
	// Markdown 类型消息
	msgTypeMarkdown = "markdown"
	// ActionCard 类型消息
	msgTypeActionCard = "actionCard"
	// FeedCard 类型消息
	msgTypeFeedCard = "feedCard"
)

// Text 类型消息
type TextMessage struct {
	// 首屏会话透出的展示内容
	Content string `json:"content"`
}

// Link 类型消息
type LinkMessage struct {
	// 首屏会话透出的展示内容
	Text string `json:"text"`
	// 消息标题
	Title string `json:"title"`
	// 图片URL
	PicUrl string `json:"picUrl"`
	// 点击消息跳转的URL
	MessageUrl string `json:"messageUrl"`
}

// Markdown 类型消息
type MarkdownMessage struct {
	// 首屏会话透出的展示内容
	Title string `json:"title"`
	// markdown格式的消息
	Text string `json:"text"`
}

// ActionCard
type ActionCard struct {
	// 首屏会话透出的展示内容
	Title string `json:"title"`
	// markdown格式的消息
	Text string `json:"text"`
	// 0-按钮竖直排列，1-按钮横向排列
	BtnOrientation string `json:"btnOrientation"`
}

type ActionCardMessage struct {
	ActionCard
	// 单个按钮的方案。(设置此项和singleURL后btns无效。)
	SingleTitle string `json:"singleTitle"`
	// 点击singleTitle按钮触发的URL
	SingleURL string `json:"singleURL"`
}

type ActionCardMessage2 struct {
	ActionCard
	// 按钮列表
	Btns []ActionCardMessageBtn `json:"btns"`
}

type ActionCardMessageBtn struct {
	// 按钮方案
	Title string `json:"title"`
	// 点击按钮触发的URL
	ActionURL string `json:"actionURL"`
}

// FeedCard
type FeedCardMessage struct {
	// 单条信息文本
	Links []FeedCardMessageLink `json:"links"`
}

type FeedCardMessageLink struct {
	// 单条信息文本
	Title string `json:"title"`
	// 单条信息后面图片的URL
	MessageURL string `json:"messageURL"`
	// 单条信息后面图片的URL
	PicURL string `json:"picURL"`
}
