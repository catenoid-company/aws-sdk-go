// +build integration

//Package elasticache provides gucumber integration tests support.
package elasticache

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/elasticache"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@elasticache", func() {
		World["client"] = elasticache.New(smoke.Session)
	})
}
