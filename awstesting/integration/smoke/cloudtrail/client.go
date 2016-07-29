// +build integration

//Package cloudtrail provides gucumber integration tests support.
package cloudtrail

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/cloudtrail"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudtrail", func() {
		World["client"] = cloudtrail.New(smoke.Session)
	})
}
