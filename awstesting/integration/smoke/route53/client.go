// +build integration

//Package route53 provides gucumber integration tests support.
package route53

import (
	"github.com/catenoid-company/aws-sdk-go/awstesting/integration/smoke"
	"github.com/catenoid-company/aws-sdk-go/service/route53"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@route53", func() {
		World["client"] = route53.New(smoke.Session)
	})
}
