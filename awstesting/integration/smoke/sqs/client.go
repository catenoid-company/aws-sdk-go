// +build integration

//Package sqs provides gucumber integration tests support.
package sqs

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/sqs"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@sqs", func() {
		World["client"] = sqs.New(smoke.Session)
	})
}
