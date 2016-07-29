// +build integration

//Package ecs provides gucumber integration tests support.
package ecs

import (
	"github.com/catenoid-company/aws-sdk-go/aws"
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/ecs"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@ecs", func() {
		// FIXME remove custom region
		World["client"] = ecs.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
