// +build integration

//Package cloudfront provides gucumber integration tests support.
package cloudfront

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/cloudfront"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudfront", func() {
		World["client"] = cloudfront.New(smoke.Session)
	})
}
