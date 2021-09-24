package awsgo

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

//定义sess调用方法 读取本地config
//sess := session.Must(session.NewSessionWithOptions(session.Options{
//	SharedConfigState: session.SharedConfigEnable,
//}))
func getalerms(sess *session.Session) (*cloudwatch.DescribeAlarmsOutput, error) {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.

	svc := cloudwatch.New(sess)

	resp, err := svc.DescribeAlarms(nil)
	if err != nil {
		fmt.Println("Got error getting alarm descriptions:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for a, alarm := range resp.MetricAlarms {
		fmt.Println(a, alarm)
	}
}
