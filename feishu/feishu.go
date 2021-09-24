package feishu

import (
	"fmt"
	"net/http"
	"strings"
)

func sendMsg(apiUrl, msg string) {
	// json
	contentType := "application/json"
	// data
	sendData := `{
		"msg_type": "text",
		"content": {"text": "` + "消息通知:" + msg + `"}
	}`
	// request
	result, err := http.Post(apiUrl, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()

}

/* func main() {
	// webhook地址
	var webhookUrl string
	// 消息内容
	var message string

	flag.StringVar(&webhookUrl, "u", "", "请输入webook地址")
	flag.StringVar(&message, "s", "", "请输入消息")

	flag.Parse()
	flag.Usage()
	sendMsg(webhookUrl, message)
}
*/
