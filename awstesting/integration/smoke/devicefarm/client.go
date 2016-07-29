// +build integration

//Package devicefarm provides gucumber integration tests support.
package devicefarm

import (
	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/devicefarm"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@devicefarm", func() {
		// FIXME remove custom region
		World["client"] = devicefarm.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
