// +build integration

//Package cloudwatchlogs provides gucumber integration tests support.
package cloudwatchlogs

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/cloudwatchlogs"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudwatchlogs", func() {
		World["client"] = cloudwatchlogs.New(smoke.Session)
	})
}
