// +build integration

//Package iotdataplane provides gucumber integration tests support.
package iotdataplane

import (
	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/iot"
	"github.com/catenoid-company/aws-sdk-go/service/iotdataplane"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@iotdataplane", func() {
		svc := iot.New(smoke.Session)
		result, err := svc.DescribeEndpoint(&iot.DescribeEndpointInput{})
		if err != nil {
			World["error"] = err
			return
		}

		World["client"] = iotdataplane.New(smoke.Session, aws.NewConfig().
			WithEndpoint(*result.EndpointAddress))
	})
}
