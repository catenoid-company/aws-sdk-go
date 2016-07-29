// +build integration

//Package lambda provides gucumber integration tests support.
package lambda

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/lambda"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@lambda", func() {
		World["client"] = lambda.New(smoke.Session)
	})
}
