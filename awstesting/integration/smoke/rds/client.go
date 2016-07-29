// +build integration

//Package rds provides gucumber integration tests support.
package rds

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/rds"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@rds", func() {
		World["client"] = rds.New(smoke.Session)
	})
}
