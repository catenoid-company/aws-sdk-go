// +build integration

//Package simpledb provides gucumber integration tests support.
package simpledb

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/simpledb"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@simpledb", func() {
		World["client"] = simpledb.New(smoke.Session)
	})
}
