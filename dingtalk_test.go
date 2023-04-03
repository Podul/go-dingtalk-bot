package dingtalk

import (
	"testing"
)

type d = struct {
	key  string
	data string
	sign string
}

func TestA(t *testing.T) {
	// rows := []d{
	// 	{key: "secret", data: "1678428938751\nsecret", sign: "aOfYRpHZt2BKzMpNFvvyYqkZteDWHIC27bSveBCP5UM="},
	// 	{key: "secret1", data: "167842893871\nsecret1", sign: "jkhN5lBgQvAe4y8qajNauM9oq9n/HEBcBTZfnqyLQS8="},
	// }

	// for _, v := range rows {
	// 	sign := sha256SignToBase64(v.key, v.data)
	// 	if v.sign == sign {
	// 		t.Logf("{key: %s, data: %s} 生成 sign: %s 与 %s 匹配", v.key, v.data, sign, v.sign)
	// 	} else {
	// 		t.Errorf("{key: %s, data: %s} 生成 sign: %s 与 %s 不匹配", v.key, v.data, sign, v.sign)
	// 	}
	// }
}
