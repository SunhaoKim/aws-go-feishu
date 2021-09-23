package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func sendMsg(msg string) {
	// json
	contentType := "application/json"
	// data
	sendData := `{
		"msg_type": "text",
		"content": {"text": "` + "消息通知:" + msg + `"}
	}`
	// request
	result, err := http.Post("webhook地址", contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()

}

func main() {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cloudwatch.New(sess)

	resp, err := svc.DescribeAlarms(nil)
	if err != nil {
		fmt.Println("Got error getting alarm descriptions:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, alarm := range resp.MetricAlarms {
		sendMsg(*alarm.AlarmName)
	}
}

