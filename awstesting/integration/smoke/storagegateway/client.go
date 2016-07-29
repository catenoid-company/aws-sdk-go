// +build integration

//Package storagegateway provides gucumber integration tests support.
package storagegateway

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/storagegateway"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@storagegateway", func() {
		World["client"] = storagegateway.New(smoke.Session)
	})
}
