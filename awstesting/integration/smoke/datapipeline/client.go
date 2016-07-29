// +build integration

//Package datapipeline provides gucumber integration tests support.
package datapipeline

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/datapipeline"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@datapipeline", func() {
		World["client"] = datapipeline.New(smoke.Session)
	})
}
