package main

import (
	"fmt"
	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/aws/session"
	"github.com/catenoid-company/aws-sdk-go/service/ec2"
	"log"
	"os"
	"strings"
)

// This example will list instances with a filter
//
// Usage:
// go run filter_ec2_by_tag.go <name_filter>
func main() {
	nameFilter := os.Args[1]
	awsRegion := "us-east-1"
	svc := ec2.New(session.New(&aws.Config{Region: aws.String(awsRegion)}))
	fmt.Printf("listing instances with tag %v in: %v\n", nameFilter, awsRegion)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", nameFilter, "*"}, "")),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", awsRegion, err.Error())
		log.Fatal(err.Error())
	}
	fmt.Printf("%+v\n", *resp)
}
