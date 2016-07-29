// +build integration

//Package cloudwatch provides gucumber integration tests support.
package cloudwatch

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/cloudwatch"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudwatch", func() {
		World["client"] = cloudwatch.New(smoke.Session)
	})
}
