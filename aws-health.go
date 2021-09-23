// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/health"
)

//func GetEvent(sess *session.Session) (*health.DescribeEventsOutput, error) {
//	svc := health.New(sess)
//	result, err := svc.DescribeEvents(nil)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, err
//}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := health.New(sess)
	// result, err := GetEvent(sess)
	result, err := svc.DescribeEvents(nil)
	if err != nil {
		fmt.Println("Get event error")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(result)

	//if err != nil {
	//	fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
	//	fmt.Println(err)
	//	return
	//}

}

