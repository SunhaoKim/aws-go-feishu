// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package awsgo

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/health"
)

//定义sess调用方法 读取本地config
//sess := session.Must(session.NewSessionWithOptions(session.Options{
//	SharedConfigState: session.SharedConfigEnable,
//}))
func gethealth(sess *session.Session) (*health.DescribeEventsOutput, error) {

	svc := health.New(sess)
	// result, err := GetEvent(sess)
	result, err := svc.DescribeEvents(nil)
	if err != nil {
		fmt.Println("Get event error")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(result)

}
