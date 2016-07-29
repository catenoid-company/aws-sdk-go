// +build integration

//Package cloudsearch provides gucumber integration tests support.
package cloudsearch

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/cloudsearch"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudsearch", func() {
		World["client"] = cloudsearch.New(smoke.Session)
	})
}
